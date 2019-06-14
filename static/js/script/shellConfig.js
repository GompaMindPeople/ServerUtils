
init()
function init() {
    console.log("init")
    $.ajax({
        url:"/listSSHCombobox",
        type:"get",
        success:function(data){
            $('#combobox_hostName').combobox({
                valueField:'SSHId',
                textField:'HostName',
                onSelect:function(record){
                    $("#SSHId").val(record.SSHId)
                    $.ajax({
                        url:"/getSshConfigById",
                        type:"get",
                        data:{SSHId:record.SSHId},
                        success:function(result){
                            $('#saveForm').form("load",result)
                        }
                    })
                }
            });
            $('#combobox_hostName').combobox("loadData",data)
        }
    })

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

 
 


 
 

