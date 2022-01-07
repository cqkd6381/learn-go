
function StringObj() {

}
StringObj.prototype.checkEmpty = function(str){
    if (str === 'undefined' || !str || !/[^\s]/.test(str)) {
        return 1;
    }

    return 0;
};
