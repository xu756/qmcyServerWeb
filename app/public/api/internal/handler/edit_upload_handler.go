package handler

import (
	"context"
	"github.com/tencentyun/cos-go-sdk-v5"
	"github.com/xu756/qmcy/app/public/api/internal/logic"
	"github.com/xu756/qmcy/app/public/api/internal/svc"
	"github.com/xu756/qmcy/common/result"
	"github.com/xu756/qmcy/common/tool"
	"github.com/xu756/qmcy/common/xerr"
	"log"
	"net/http"
)

func EditUploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		file, _, err := r.FormFile("file")
		if err != nil {
			result.HttpResult(r, w, nil, xerr.NewMsgError("无文件"))
			return
		}

		name := tool.NewUid() + ".jpg"
		defer file.Close()

		_, err = svcCtx.CosClient.Object.Put(context.Background(), "images/"+name, file, nil)
		if err != nil {
			result.HttpResult(r, w, nil, xerr.NewMsgError("上传失败"+err.Error()))
			return
		}

		pic := &cos.PicOperations{
			IsPicInfo: 1,
			Rules: []cos.PicOperationsRules{
				{
					FileId: name,
					Rule:   "imageMogr2/thumbnail/1125x/format/jpg/interlace/1/quality/73/strip",
				},
			},
		}
		opt := &cos.ObjectPutOptions{
			ObjectPutHeaderOptions: &cos.ObjectPutHeaderOptions{
				XOptionHeader: &http.Header{},
			},
		}
		opt.ObjectPutHeaderOptions.XOptionHeader.Set("Pic-Operations", cos.EncodePicOperations(pic))
		res, _, err := svcCtx.CosClient.CI.ImageProcess(context.Background(), "images/"+name, pic)
		if err != nil {
			result.HttpResult(r, w, nil, xerr.NewMsgError("压缩图片失败"+err.Error()))
			return
		}
		log.Print(res)
		l := logic.NewEditUploadLogic(r.Context(), svcCtx)
		resp, err := l.EditUpload(name)
		result.HttpResult(r, w, resp, err)
	}

}
