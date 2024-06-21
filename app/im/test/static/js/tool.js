var gloablLoading
function loading(text){
    if (text){
        gloablLoading = ELEMENT.Loading.service({text: text,background:"rgba(0, 0, 0, 0.8)"})
    } else {
        gloablLoading.close()
    }
}
function toast(text){
    ELEMENT.Message({
        message: text,
    });
}
function toastSucc(text){
    ELEMENT.Message({
        message: text,
        type: 'success'
    });
}
function toastError(text){
    ELEMENT.Message.error(text);
}
var tool = {
    loading,
    toast,
    toastSucc,
    toastError
}