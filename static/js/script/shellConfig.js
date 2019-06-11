
init()
function init() {
    console.log("init")
    $.ajax({
        url:"/getSshConfig",
        type:"get",
        success:function(data){
            console.log(data)
            if (data) {
                $("#userName").val(data.UserName)
                $("#password").val(data.Password)
                $("#hostName").val(data.HostName)
                $("#port").val(data.Port)
            }
        }
    })

    $('#submitBtn').bind('click', function(){
        saveConfig()
    });
    $('#btn').linkbutton({
        iconCls: 'icon-search'
    });

}

function saveConfig() {
    $('#saveForm').form('submit', {
        url: "/saveSSHConfig",
        onSubmit: function () {
            var isValid = $(this).form('validate');
            if (!isValid) {

            }
            fileData = $(this).serialize();
            return isValid;
        },
        success: function (result) {
            alert("保存成功")

        }
    });
}

 
 


 
 

