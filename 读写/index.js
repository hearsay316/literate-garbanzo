var name = "The Window";

var object = {
    name : "My Object",
    time:new Date(),
    getNameFunc : function(){
        return function(){
            return this

        };

    }

};

console.log(object.time);
