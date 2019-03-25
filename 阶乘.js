var sum = 1;
function main(num) {
    if(typeof num==="number"){
       if(num===1){
           return ;
       }
        main(num-1);
        sum*=num;
    }else {
        return new Error("不是数字");
    }

}
main(4);
console.log(sum);