// Initialize your app
var myApp = new Framework7();

// Export selectors engine
var $$ = Dom7;

// Add view
var mainView = myApp.addView('.view-main', {
    // Because we use fixed-through navbar we can enable dynamic navbar
    dynamicNavbar: true
});

// Callbacks to run specific code for specific pages, for example for About page:
myApp.onPageInit('about', function(page) {
    // run createContentPage func after link was clicked
    $$('.create-page').on('click', function() {
        createContentPage();
    });
});

// Generate dynamic page
var dynamicPageIndex = 0;

function createContentPage() {
    mainView.router.loadContent(
        '<!-- Top Navbar-->' +
        '<div class="navbar">' +
        '  <div class="navbar-inner">' +
        '    <div class="left"><a href="#" class="back link"><i class="icon icon-back"></i><span>Back</span></a></div>' +
        '    <div class="center sliding">Dynamic Page ' + (++dynamicPageIndex) + '</div>' +
        '  </div>' +
        '</div>' +
        '<div class="pages">' +
        '  <!-- Page, data-page contains page name-->' +
        '  <div data-page="dynamic-pages" class="page">' +
        '    <!-- Scrollable page content-->' +
        '    <div class="page-content">' +
        '      <div class="content-block">' +
        '        <div class="content-block-inner">' +
        '          <p>Here is a dynamic page created on ' + new Date() + ' !</p>' +
        '          <p>Go <a href="#" class="back">back</a> or go to <a href="services.html">Services</a>.</p>' +
        '        </div>' +
        '      </div>' +
        '    </div>' +
        '  </div>' +
        '</div>'
    );
    return;
}

//-------------------------------------------------------------

// 加载flag
var loading = false;

// 上次加载的序号
var lastIndex = $$('.list-block li').length;

// 最多可加载的条目
var maxItems = 100;

// 注册'infinite'事件处理函数
$$('.infinite-scroll').on('infinite', function() {
    // 如果正在加载，则退出
    if (loading) return;
    // 设置flag
    loading = true;
    // 模拟1s的加载过程
    setTimeout(function() {
        // 重置加载flag
        loading = false;
        if (lastIndex >= maxItems) {
            // 加载完毕，则注销无限加载事件，以防不必要的加载
            myApp.detachInfiniteScroll($$('.infinite-scroll'));
            // 删除加载提示符
            $$('.infinite-scroll-preloader').remove();
            return;
        }

        var html = '';

        //获取新记录
        var xmlhttp = new XMLHttpRequest();
        xmlhttp.onreadystatechange = function() {
            if (xmlhttp.readyState == 4 && xmlhttp.status == 200) {
                records = JSON.parse(xmlhttp.responseText);
                for (i in records) {
                    html += '<li>' +
                        '<a href="views/details.html" class="item-link">' +
                        '<div class="item-content">' +
                        '<div class="item-media">' +
                        '<img src="' + records[i].ItempicPath +'" height="80"></div>' +
                        '<div class="item-inner">' +
                        '<div class="item-title-row">' +
                        '<div class="item-title">' + records[i].Title +'</div>' +
                        '<div class="item-after">' + records[i].Date +'</div>' +
                        '</div>' +
                        '<div class="item-subtitle">' + records[i].Category +'</div>' +
                        '<div class="item-text">' + records[i].Summary +'</div>' +
                        '</div>' +
                        '</div>' +
                        '</a>' +
                        '</li>'
                }
                // 添加新条目
                $$('.list-block ul').append(html);

                // 更新最后加载的序号
                lastIndex = $$('.list-block li').length;
            }
        }
        xmlhttp.open("GET", "/api/list/" + lastIndex, true);
        xmlhttp.send();

    }, 1000);
});

//-------------------------------------------------------------