require(["jquery","jquery.scrollTo","vue"],function($ , _ , Vue){

// 背景
    $("#firstStep").click(function(){
        $.scrollTo('#vue',500);
    });

    Vue.config.devtools = true
    new Vue({
        el : "#vue",
        data : {
            lists : JSON.parse(MODEL) ,
        },
        methods : {

        },
        mounted: function(){
            console.log(this.lists)
        },
    });
});