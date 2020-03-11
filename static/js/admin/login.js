layui.use(['form'], function () {
    var form = layui.form,
        layer = layui.layer;

    // 登录过期的时候，跳出ifram框架
    if (top.location != self.location) top.location = self.location;

    // 粒子线条背景
    $(document).ready(function () {
        $('.layui-container').particleground({
            dotColor: '#5cbdaa',
            lineColor: '#5cbdaa'
        });
    });

    // 进行登录操作
    form.on('submit(login)', function (data) {
        data = data.field;
        $.ajax({
            url: 'login',
            type: 'post',
            contentType: "application/x-www-form-urlencoded; charset=UTF-8",
            dataType: "json",
            data: data,
            success: function (res) {
                alert('请求成功')
            },
            error: function (xhr, textstatus, thrown) {
                layer.msg('网络异常')
                return false;
            }
        });
        return false;
    });
});