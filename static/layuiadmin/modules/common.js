/** layuiAdmin.std-v1.1.0 LPPL License By http://www.layui.com/admin/ */
;layui.define(function (e) {
    var i = (layui.$, layui.layer, layui.laytpl, layui.setter, layui.view, layui.admin);
    i.events.logout = function () {
        location.href = "/admin/logout";
        // i.req({
        //     url: "/admin/logout", type: "get", data: {}, done: function (e) {
        //         i.exit(function () {
        //             location.href = "/admin/login"
        //         })
        //     }
        // })
    }, e("common", {})
});