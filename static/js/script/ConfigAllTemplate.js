
init()
function init() {
    console.log("init")
    $.ajax({
        url:"/listServerName",
        type:"get",
        success:function(data){
            console.log(data)
            if (data) {
                $('#combobox_serverName').combobox({
                    valueField:'TemplateId',
                    textField:'ServerName',
                    onSelect:function(record){
                        $("#TemplateId").val(record.TemplateId)
                        $.ajax({
                            url:"/getTemplateById",
                            type:"get",
                            data:{TemplateId:record.TemplateId},
                            success:function(result){
                                $('#saveForm').form("load",result)
                            }
                        })
                    }
                });


                $("#combobox_serverName").combobox("loadData", data);
            }
        }
    })

    $('#btn').linkbutton({
        iconCls: 'icon-search'
    });

}

function downloadConfigAllTemplate() {
    //generateConfigAllTemplate
   let template = $("#TemplateId").val()
    $.ajax({
        url:"/generateConfigAllTemplate",
        data:{templateId:template},
        type:"get",
        success:function (data) {
            window.open('http://localhost:8081/downloadConfigAllTemplate')
        }
    })
}


function saveConfigAllTemplate() {
    $('#saveForm').form('submit', {
        url: "/saveConfigAll",
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

 
 


 
 

