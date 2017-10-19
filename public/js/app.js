
require.config({
    baseUrl: "/js/",
    paths: {
        "jquery": "//cdn.bootcss.com/jquery/2.2.1/jquery.min",
        "vue" : "//cdn.bootcss.com/vue/2.4.4/vue",
        "jquery.scrollTo" : "//cdn.bootcss.com/jquery-scrollTo/2.1.2/jquery.scrollTo.min" ,
        "index" : "index",
        "videojs" : "//cdn.bootcss.com/video.js/6.3.3/video.min"
    },
    shim:{
        vue:{exports:'Vue'} //注意这行
    }
});
