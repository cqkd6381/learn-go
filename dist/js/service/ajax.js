function fetchJson(headers, url, method, data)
{
    $.ajax({
        // headers: {
        //     'X-auth-token': xAuthToken
        // },
        // url: rec_api + 'api/v2/publicFileUpload',
        // type: 'POST',
        // cache: false, //上传文件不需要缓存
        // data: formData,
        // processData: false, // 告诉jQuery不要去处理发送的数据
        // contentType: false, // 告诉jQuery不要去设置Content-Type请求头
        // success: function (data) {
        //     let object_id = data.entity.object_id;
        //     document.getElementById("my-img").src = data.entity.file_url;
        //     $('#modal_cloud_file_id').val(object_id);
        // },
        // error: function (data) {
        //     message({
        //         type: 'error',
        //         message: res.message
        //     })
        // }


    })
}
