function quObject(obj1,obj2) {
    if(JSON.stringify(obj1).length !==JSON.stringify(obj2).length){
        return false
    }
    if(typeof obj1 !=="object"&&typeof obj2 !=="object"){
        return false
    }
    if (obj1.length||obj2.length) {
        return false;
    }
    for (let a in obj1){
        if(obj2.a&&obj2.a===obj1.a){
            return true;
        }else {
            return false;
        }
    }
}
quObject({a:10},{b:20});
