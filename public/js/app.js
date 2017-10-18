
require.config({
    baseUrl: "/js/",
    paths: {
        "jquery": "//cdn.bootcss.com/jquery/2.2.1/jquery.min",
        "vue" : "//cdn.bootcss.com/vue/2.4.4/vue",
        "jquery.scrollTo" : "//cdn.bootcss.com/jquery-scrollTo/2.1.2/jquery.scrollTo.min" ,
        "index" : "index"
    },
    shim:{
        vue:{exports:'Vue'} //注意这行
    }
});
