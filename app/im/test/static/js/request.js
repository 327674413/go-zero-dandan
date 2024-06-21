const requestUrl = "http://localhost:8088"
function request(obj) {
    return new Promise((resolve, reject)=>{
        let congtroller = obj.c;
        let method = obj.method ? obj.method : 'POST';
        let headers = obj.headers ? obj.headers : {};
        if (!obj.data) obj.data = {};
        if (obj.loading) tool.loading(obj.loading);
        let url = obj.url ? obj.url : requestUrl+ '/'+congtroller ;
        let responseType = obj.responseType ? obj.responseType : '';
        return axios({
            url: url,
            data: obj.data,
            method: method,
            headers:headers,
            responseType:responseType
        }).then(res=>{
            if (obj.loading)  tool.loading();
            if (responseType == 'blob'){
                resolve(res);
            }else if (res.data.result){
                resolve(res.data.data);
            } else if (res.data.result === false) {
                if(obj.reject){
                    reject(res.data);
                } else {
                    tool.toastError(res.data.msg)
                }
            } else {
                if (obj.reject) {
                    reject(res);
                } else{
                    tool.toastError('遇到了点小问题');
                }
            }
        }).catch(error=>{
            if (obj.loading)  tool.loading()
            if (obj.reject){
                reject(error)
            } else {
                tool.toastError('请求失败');
            }

        });
    });

}