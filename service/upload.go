package service

// UploadToQiNiu 七牛云上传
// func UploadToQiNiu(file multipart.File, fileSize int64) (path string, err error) {
//     var AccessKey = conf.AccessKey
//     var SerectKey = conf.SerectKey
//     var Bucket = conf.Bucket
//     var ImgUrl = conf.QiniuServer
//     putPlicy := storage.PutPolicy{
//         Scope: Bucket,
//     }
//     mac := qbox.NewMac(AccessKey, SerectKey)
//     upToken := putPlicy.UploadToken(mac)

//     cfg := storage.Config{
//         Zone: &storage.ZoneHuanan,  // 地区
//         UseCdnDomains: false,       // CDN
//         UseHTTPS: false,            // 不使用https
//     }

//     formUploader := storage.NewFormUploader(&cfg)
//     putExtra := storage.PutExtra{}

//     ret := storage.PutRet{}
//     err = formUploader.PutWithoutKey(context.Background(), &ret, upToken, file, fileSize, &putExtra)
//     if err != nil {
//         return "", err
//     }
//     url := ImgUrl + ret.Key
//     return url, nil
// }