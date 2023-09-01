package handler

import (
	"context"
	"github.com/tencentyun/cos-go-sdk-v5"
	"github.com/xu756/qmcy/common/result"
	"github.com/xu756/qmcy/common/tool"
	"github.com/xu756/qmcy/common/xerr"
	"net/http"

	"github.com/xu756/qmcy/app/public/api/internal/logic"
	"github.com/xu756/qmcy/app/public/api/internal/svc"
)

func UploadOneImgHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		file, _, err := r.FormFile("file")
		if err != nil {
			result.HttpResult(r, w, nil, xerr.NewMsgError("无文件"))
			return
		}
		path := r.FormValue("path")
		name := tool.NewUid() + ".jpg"
		defer file.Close()

		_, err = svcCtx.CosClient.Object.Put(context.Background(), path+"/"+name, file, nil)
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
		_, _, err = svcCtx.CosClient.CI.ImageProcess(context.Background(), path+"/"+name, pic)
		if err != nil {
			result.HttpResult(r, w, nil, xerr.NewMsgError("压缩图片失败"+err.Error()))
			return
		}

		l := logic.NewUploadOneImgLogic(r.Context(), svcCtx)
		resp, err := l.UploadOneImg(path + "/" + name)
		result.HttpResult(r, w, resp, err)
	}
}
