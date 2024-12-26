package logic

import (
	"context"
	"github.com/ac-dcz/lightRW/apps/code/rpc/verifycode"
	"github.com/ac-dcz/lightRW/apps/user/model"
	"github.com/ac-dcz/lightRW/apps/user/rpc/internal/svc"
	"github.com/ac-dcz/lightRW/apps/user/rpc/pb"
	"github.com/ac-dcz/lightRW/common/codes"
	"github.com/ac-dcz/lightRW/common/errors"
	"github.com/ac-dcz/lightRW/common/utils"

	stderr "errors"
	"github.com/zeromicro/go-zero/core/logx"
)

const passWordPrefix = "password#"

type SignUpLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSignUpLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SignUpLogic {
	return &SignUpLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SignUpLogic) SignUp(in *pb.SignUpReq) (*pb.SignUpResp, error) {
	//Step0: 参数的有效性 由api端检查
	//Step1: 验证Code
	if resp, err := l.svcCtx.CodeRpc.VerifyCode(l.ctx, &verifycode.VerifyCodeReq{
		Code: in.Code,
		Tel:  in.Tel,
	}); err != nil {
		l.Logger.Errorf("VerifyCode err:%v", err)
		return nil, err
	} else if !resp.Success {
		l.Logger.Error("verify code failed")
		return nil, errors.New(codes.VerifyCodeErr, "verify code error")
	}

	//Step2: 检查Tel是否已经注册
	if _, err := l.svcCtx.UserModel.FindOneByTel(l.ctx, in.Tel); !stderr.Is(err, model.ErrNotFound) {
		if err != nil {
			l.Logger.Errorf("FindOneByTel err:%v", err)
			return nil, errors.New(codes.InternalError, err.Error())
		}
		return nil, errors.New(codes.TelAlreadyExists, "tel is already exists")
	}

	//Step3: 加密密码
	pwd := utils.MD5(passWordPrefix + in.Pass)
	user := &model.User{
		NickName: in.NickName,
		Password: pwd,
		Tel:      in.Tel,
		Level:    []byte{byte(in.Level)},
		Status:   model.StatusOk,
	}

	//Step4: 更新数据库
	if r, err := l.svcCtx.UserModel.Insert(l.ctx, user); err != nil {
		l.Logger.Errorf("InsertUser err:%v", err)
		return nil, errors.New(codes.InternalError, err.Error())
	} else {
		id, _ := r.LastInsertId()
		return &pb.SignUpResp{
			User: &pb.User{
				NickName: in.NickName,
				Tel:      in.Tel,
				Level:    in.Level,
				Uid:      uint64(id),
			},
		}, nil
	}
}
