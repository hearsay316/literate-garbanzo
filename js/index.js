function num(s) {
    let len = s.length
    let str = "";
    let arr = [];
    if (len >= 1204) {
        return new Error("错误")
    }
    for (let i = 0; i < len; i++) {
        if ("0123456789.".indexOf(s[i]) != -1) {
            if (str[str.length - 1] == "." && s[i] == ".") {
                str = ""
            } else {
                if (str.indexOf(".") > 0) {
                    str += s[i]
                    let a = str.split(".")
                    if (a[1].length >= a[0].length) {
                        arr.push(a[1])
                    } else {
                        arr.push(str)
                    }
                } else {
                    str += s[i]
                    arr.push(str)
                }
            }
        } else {
            str = ""
        }
    }
    console.log(arr);
    return arr.sort((a, b) => b - a)[0]
}

function objc(s) {
    let len = s.length;
    let a = 0;
    let obj = {};
    if (len >= 1024) {
        return
    }
    for (let i = 0; i < len; i++) {
        if ("0123456789.".indexOf(s[i]) != -1) {
            if (obj[a]) {
                if(obj[a].indexOf(".") >0 && s[i] === "."){
                    a++
                }else {
                    if(obj[a].length>0&&s[i]!==".") {
                        if (obj[a].indexOf(".") > 0) {
                            let b = (obj[a] + s[i]).split(".");
                            if (b[1].length >= b[0].length) {
                                if(s[i+1]!=="."){
                                    obj[a] = b[1];
                                }else {
                                    obj[a] = b[1];
                                    a++
                                }
                            } else {
                                obj[a] += s[i];

                            }
                        } else {
                            obj[a] += s[i];

                        }
                    }else {
                            obj[a] += s[i];

                    }
                }
            } else {
                if(s[i]!=="."){
                    obj[a] = s[i]
                }
            }
        } else {
            a++
        }
    }
    console.log(Object.values(obj));
    return Object.values(obj).sort((q,w)=>w-q)[0]
}

/*let a = num("54.4..6ll-44.4++5-4")
let b = num("ppp12bbbbb12.31ccc")
let c = num("asc23423v234")*/
let a = objc("54.4..6ll-44.4++5-4");
let b = objc("ppp12bbbbb12.31ccc");
let c = objc("asc23423v234");
let d = objc("1.8.5.9.2.1.9");

console.log(a,c,b,d);