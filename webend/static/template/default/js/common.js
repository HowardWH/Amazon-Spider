/**通用的JS**/
//通用返回顶部
$(function(){
    //首先将#back-to-top隐藏
    $("#slider-goTop").hide();
    //当滚动条的位置处于距顶部100像素以下时，跳转链接出现，否则消失
    $(window).scroll(function(){
        if ($(window).scrollTop()>100){
            $("#slider-goTop").fadeIn();
        }else{
            $("#slider-goTop").fadeOut();
        }
    });
    //当点击跳转链接后，回到页面顶部位置
    $("#slider-goTop").click(function(){
        $('body,html').animate({scrollTop:0},500);
        return false;
    });
    //返回顶部等滑块hover事件
    $('#slider-chat,#slider-qq,#slider-phone,#slider-wechat').hover(
        function(){
            $(this).next().show();
        },
        function(){
            $(this).next().hide();
        }
    );
});


//收藏本站
function AddFavorite(title, url) {
    try {
        window.external.addFavorite(url, title);
    }
    catch (e) {
        try {
            window.sidebar.addPanel(title, url, "");
        }
        catch (e) {
            alert("抱歉，您所使用的浏览器无法完成此操作。\n\n加入收藏失败，请使用Ctrl+D进行添加");
        }
    }
}
