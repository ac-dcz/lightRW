package logic

import (
	"context"
	"fmt"
	"github.com/ac-dcz/lightRW/apps/code/rpc/internal/svc"
	"github.com/ac-dcz/lightRW/apps/code/rpc/pb"
	"github.com/ac-dcz/lightRW/common/codes"
	"github.com/ac-dcz/lightRW/common/errors"
	"github.com/ac-dcz/lightRW/common/utils"
	"strconv"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

const (
	RdsCodePrefix         = "biz:verify_code:"
	RdsCodeNumsPrefix     = "biz:verify_code:nums:"
	RdsCodeIntervalPrefix = "biz:verify_code:interval:"
)

func codePrefix(tel string) string {
	return fmt.Sprintf("%s%s", RdsCodePrefix, tel)
}

func codeNumsPrefix(tel string) string {
	return fmt.Sprintf("%s%s", RdsCodeNumsPrefix, tel)
}

func codeIntervalPrefix(tel string) string {
	return fmt.Sprintf("%s%s", RdsCodeIntervalPrefix, tel)
}

const (
	ExpiredTime = time.Minute * 5 // 验证码的过期时间
	MaxNums     = 10              // 1h 内最大申请验证码次数
	NumsTime    = time.Hour
	Interval    = time.Second * 60 //两次验证码之间的最小间隔
)

type GenCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGenCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenCodeLogic {
	return &GenCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GenCodeLogic) GenCode(in *pb.GenCodeReq) (*pb.GenCodeResp, error) {
	//Step0: in.Tel 的有效性由API层检查
	//Step1: 检查调用间隔
	intervalKey := codeIntervalPrefix(in.Tel)
	if ok, err := l.svcCtx.BizRds.SetnxExCtx(l.ctx, intervalKey, "", int(Interval.Seconds())); err != nil {
		l.Logger.Errorf("generator verify code error: %v", err.Error())
		return nil, errors.New(codes.InternalError, err.Error())
	} else if !ok {
		l.Logger.Infow("generator verify code too fast", logx.Field("tel", in.Tel))
		return nil, errors.New(codes.GenCodeFast, "generator verify code too fast")
	}

	//Step2: 检查是否超过最大调用次数
	numsKey := codeNumsPrefix(in.Tel)
	if nums, err := l.getGenCodeNums(numsKey); err != nil {
		l.Logger.Errorf("generator verify code error: %v", err.Error())
		return nil, errors.New(codes.InternalError, err.Error())
	} else if nums >= MaxNums {
		l.Logger.Errorf("generator verify code too many", logx.Field("tel", in.Tel))
		return nil, errors.New(codes.GenCodeMany, "generator verify code too many")
	} else {
		//次数+1
		if _, err := l.svcCtx.BizRds.IncrCtx(l.ctx, numsKey); err != nil {
			l.Logger.Errorf("generator verify code error: %v", err.Error())
			return nil, errors.New(codes.InternalError, err.Error())
		}
	}

	//Step3: 生成验证码
	code := utils.GenCode(l.svcCtx.Config.CodeLen)
	codeKey := codePrefix(in.Tel)
	if err := l.svcCtx.BizRds.SetexCtx(l.ctx, codeKey, code, int(ExpiredTime.Seconds())); err != nil {
		l.Logger.Errorf("generator verify code error: %v", err.Error())
		return nil, errors.New(codes.InternalError, err.Error())
	}

	return &pb.GenCodeResp{
		Code: code,
	}, nil
}

func (l *GenCodeLogic) getGenCodeNums(key string) (int, error) {
	if _, err := l.svcCtx.BizRds.SetnxExCtx(l.ctx, key, "0", int(NumsTime.Seconds())); err != nil {
		return 0, err
	}
	if nums, err := l.svcCtx.BizRds.GetCtx(l.ctx, key); err != nil {
		return 0, err
	} else {
		n, _ := strconv.Atoi(nums)
		return n, nil
	}
}
