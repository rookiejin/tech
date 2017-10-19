require(["jquery","vue"],function($ , Vue){
    Vue.config.devtools = true;
    new Vue({
        el : "#vue",
        data : {
            series : MODEL.series ,
            videos : MODEL.videos
        },
        methods : {

        },
        mounted: function(){
            console.log(this.lists)
        },
    });
});