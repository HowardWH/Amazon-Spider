/*!
 * HoverClock(jQuery Plugin)
 * version: 2.1.0
 * Copyright (c) 2016 HoverTree
 * http://hovertree.com
 * http://hovertree.com/texiao/hoverclock/
 */
(function ($) {
    $.fn.hoverclock = function (options) {

        var settings = $.extend({
            h_radius: "10%",//边框圆角弧度
            h_borderColor: "transparent",//边框颜色
            h_backColor: "transparent",//背景演示
            h_width: "300px",//宽度
            h_height: "300px",//高度
            h_secondHandColor: "red",//秒针演示
            h_frontColor: "darkgreen",
            h_thinHandColor: "green",//分针颜色
            h_linkText: "",
            h_linkUrl: "https://www.github.com/hunterhug",
            h_linkColor: "deeppink",
            h_linkSize: "16px",
            h_hourLength: "",//时针长度
            h_minuteLength: "",//分针长度
            h_secondLength: "",//秒针长度
            h_hourNumShow: true,//是否显示正点时间数字
            h_hourNumSize: "",//正点时间数字大小
            h_hourNumColor: "green",
            h_hourNumRadii: "",//正点时间数字半径
            h_minuteNumShow: true,//是否显示逢5分钟数字
            h_minuteHeight: "",//普通分针刻度高度
            h_minuteWidth: "",
            h_minute5Height: "",//逢5分针刻度高度
            h_minute5Width: "",
            h_minute15Height: "",//逢15分针刻度高度
            h_minute15Width: ""
        }, options);

        var h_hoverClock = $(this);

        if (h_hoverClock.length < 1)
            return;
        h_hoverClock.attr('class', 'hoverclock')
        h_hoverClock.append('<div class="smallhovertree"></div><div class="bighovertree"></div><div class="secondhovertree"></div><div class="secondhovertree2"></div><div class="minutehovertree"></div><div class="hourhovertree"></div>	')
        h_hoverClock.css({ "width": settings.h_width, "height": settings.h_height, "background-color": settings.h_backColor, "border-color": settings.h_borderColor })


        var h_hoverSecond = h_hoverClock.find(".secondhovertree"), h_hoverBig = h_hoverClock.find(".bighovertree")
            , h_hoverSmall = h_hoverClock.find(".smallhovertree"), h_hoverMinute = h_hoverClock.find(".minutehovertree"), h_hoverHour = h_hoverClock.find(".hourhovertree")
        h_hoverSecond2 = h_hoverClock.find(".secondhovertree2")
        var h_minBanjing = Math.min(h_hoverClock.width(), h_hoverClock.height()) / 2

        //------------------------------------------------------
        if (settings.h_minuteWidth == "") {
            settings.h_minuteWidth = 1;
        }
        if (settings.h_minuteHeight == "") {
            settings.h_minuteHeight = h_minBanjing / 25;
        }

        if (settings.h_minute5Height == "") {
            settings.h_minute5Height = settings.h_minuteHeight + 4
        }
        if (settings.h_minute5Width == "") {
            settings.h_minute5Width = settings.h_minuteWidth + 7
        }

        if (settings.h_minute15Height == "") {

            settings.h_minute15Height = settings.h_minute5Height + 8
        }
        if (settings.h_minute15Width == "") {
            settings.h_minute15Width = settings.h_minute5Width + 4
        }



        if (settings.h_hourNumRadii == "") {
            settings.h_hourNumRadii = h_minBanjing * 4 / 5;
        }

        if (settings.h_hourNumSize == "") {
            settings.h_hourNumSize = h_minBanjing / 5;
        }
        if (settings.h_hourLength == "") {
            settings.h_hourLength = h_minBanjing / 2;
        }

        h_hoverHour.height(settings.h_hourLength)


        if (settings.h_minuteLength == "") {
            settings.h_minuteLength = h_minBanjing * 4 / 5;
        }

        h_hoverMinute.height(settings.h_minuteLength)

        if (settings.h_secondLength == "") {
            settings.h_secondLength = h_minBanjing * 24 / 25;
        }

        h_hoverSecond.height(settings.h_secondLength)
        h_hoverSecond2.height(settings.h_secondLength / 4)
        //------------------------------------------------------

        var h_now, h_second, h_minute, h_hour, h_secondRadian, h_minuteRadian, h_hourRadian;

        setInterval(
            function () {
                h_now = new Date();
                h_second = h_now.getSeconds();
                h_secondRadian = h_second * 6;
                h_minute = h_now.getMinutes();
                h_minuteRadian = ((h_minute) * 6) + (h_second / 10);
                h_hour = h_now.getHours();
                h_hourRadian = h_hour * 30 + h_minuteRadian / 12
                document.querySelector(".secondhovertree").style.transform = "rotate(" + h_secondRadian + "deg)";
                h_hoverMinute.css({ "transform": " rotate(" + h_minuteRadian + "deg)" })
                h_hoverHour.css({ "transform": " rotate(" + h_hourRadian + "deg)" })
                h_hoverSecond2.css({ "transform": " rotate(" + (h_secondRadian + 180) + "deg)" })
            }
            , 50)

        //设置圆心大小圆位置
        h_hoverBig.css({ "left": (h_hoverClock.width() - h_hoverBig.width()) / 2, "top": (h_hoverClock.height() - h_hoverBig.height()) / 2 })
        h_hoverSmall.css({ "left": (h_hoverClock.width() - h_hoverSmall.width()) / 2, "top": (h_hoverClock.height() - h_hoverSmall.height()) / 2 })

        //设置秒针位置
        h_hoverSecond.css({ "left": (h_hoverClock.width() - h_hoverSecond.width()) / 2, "top": (h_hoverClock.height() / 2 - h_hoverSecond.height()),"background-color":settings.h_secondHandColor })
        h_hoverSecond2.css({ "left": (h_hoverClock.width() - h_hoverSecond.width()) / 2, "top": (h_hoverClock.height() / 2 - h_hoverSecond2.height()), "background-color": settings.h_secondHandColor })
        //设置分针位置，边框圆角
        h_hoverMinute.css({
            "left": (h_hoverClock.width() - h_hoverMinute.width()) / 2, "top": (h_hoverClock.height() / 2 - h_hoverMinute.height())
            , "border-radius": h_hoverMinute.width() / 2
            , "background-color": settings.h_thinHandColor
        })
        //设置时针位置，边框圆角
        h_hoverHour.css({
            "left": (h_hoverClock.width() - h_hoverHour.width()) / 2, "top": (h_hoverClock.height() / 2 - h_hoverHour.height())
            , "border-radius": h_hoverHour.width() / 2
            , "background-color": settings.h_frontColor
        })

        function createsSaleHovertree(width, height, deg, color, radii) {
            var h_scale = $('<div class="scalehovertree"></div>');

            h_scale.css({ "width": width, "height": height, "z-index": 80, "transform-origin": "50% 100%", "background-color": color })
            h_scale.appendTo(h_hoverClock).css({
                "left": (h_hoverClock.width() - h_scale.width()) / 2, "top": (h_hoverClock.height() / 2 - h_scale.height()),
                "transform": "rotate(" + deg + "deg) translateY(" + (h_scale.height() - radii) + "px)"
            })
        }

        function createNum(deg, zindex, color, text, textcolor, textsize, radii) {
            var h_width = textsize * 2 + 2;
            var h_height = textsize;

            var h_scale = $('<div class="scalehovertree"></div>')

            if (text != "") {
                h_scale.html('<span style="color:' + textcolor + ';display:block;font-size:' + textsize + 'px;position:absolute;bottom:0px;font-weight:bold;text-align:center;width:100%;transform:rotate(-' + deg + 'deg)">' + text + "<span>");

            }


            h_scale.css({ "width": h_width, "height": h_height, "z-index": zindex, "transform-origin": "50% 100%", "background-color": color })
            h_scale.appendTo(h_hoverClock).css({
                "left": (h_hoverClock.width() - h_scale.width()) / 2
                , "top": h_minBanjing - h_scale.height()
                , "transform": "rotate(" + deg + "deg) translateY(" + (h_scale.height() - radii) + "px) "
            })


        }
        if (settings.h_hourNumShow) {
            //时针刻度数字
            for (var i = 0; i < 12; i++) {
                var h_deg = 360 * i / 12
                var h_hourNum = i;
                if (h_hourNum == 0)
                    h_hourNum = 12
                createNum(h_deg, 90, "transparent", h_hourNum, settings.h_hourNumColor, settings.h_hourNumSize, settings.h_hourNumRadii);
            }
        }


        var h_scaleRadii = h_minBanjing * 39 / 40


        //分针刻度数字
        for (var i = 0; i < 60; i++) {
            var h_hudu = 360 * i / 60
            var h_text = "";
            if ((i % 5) == 0) {

                if (i > 0) { h_text = i; }
                else { h_text = 60; }

                if (settings.h_minuteNumShow) {
                    createNum(h_hudu, 80, "transparent", h_text, "green", settings.h_minute5Height + 4, h_scaleRadii - 4);
                }
                else {
                    if ((i % 15) == 0) {

                        createsSaleHovertree(settings.h_minute15Width, settings.h_minute15Height, h_hudu, settings.h_frontColor, h_scaleRadii);
                    }
                    else {
                        createsSaleHovertree(settings.h_minute5Width, settings.h_minute5Height, h_hudu, settings.h_frontColor, h_scaleRadii);
                    }
                }
            }
            else {
                createsSaleHovertree(settings.h_minuteWidth, settings.h_minuteHeight, h_hudu, settings.h_frontColor, h_scaleRadii);
            }

        }


        //---------------------------------------------------
        h_hoverClock.append('<div class="hoverclocktext"><a href="https://www.github.com/hunterhug" target="_blank" class="hoverclocklink" style="text-decoration:none">珍惜时间</a></div>')
        h_hoverClock.find(".hoverclocktext").css({ "width": h_hoverClock.width(), "position": "absolute", "top": h_minBanjing * 6 / 5, "text-align": "center" })
        var h_hoverClockLink = h_hoverClock.find(".hoverclocklink");
        h_hoverClockLink.text(settings.h_linkText);
        h_hoverClockLink.attr("href", settings.h_linkUrl);
        h_hoverClockLink.attr("title", settings.h_linkText);
        h_hoverClockLink.css({ "color": settings.h_linkColor, "font-size": settings.h_linkSize })
        //----------------------------------------------------

        h_hoverClock.show();
    }
}(jQuery));