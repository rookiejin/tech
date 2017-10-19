require(["jquery","vue", "videojs"],function($ , Vue , videojs){

    new Vue({
        el:"#vue",
        data : {
          video : MODEL ,
        },
        mounted: function () {
            this.init();
        },
        methods : {
            init : function() {
                // 高度 .
                videojs('#video' , {
                    techOrder : ["html5"]
                },function () {
                    console.log('started')
                });
            }
        }
    })


});