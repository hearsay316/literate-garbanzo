function quObject(obj1, obj2) {
    let bool = true;
    console.time("b");
   /* if (JSON.stringify(obj1).length !== JSON.stringify(obj2).length) {
        return false
    }
    if (typeof obj1 !== "object" && typeof obj2 !== "object") {
        return false
    }

    if (obj1.length || obj2.length) {
        return false;
    }*/

    /*for (let a in obj1) {
        if (!obj2[a] || obj2[a] !== obj1[a]) {
            bool = false;
        }
    }*/
    // Object.keys(obj1).map((a)=>{
    //     if (!obj2[a] || obj2[a] !== obj1[a]) {
    //         bool = false;
    //     }
    // });
    console.timeEnd("b");
    return bool
}
function add(a,b){
    return a+b
}
console.time("a");

console.log(1);
console.timeEnd("a");

