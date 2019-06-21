$(function () {
    init()
    function init() {
        $('#tab').datagrid({
            url:'/listForButtonConf',
            fit:false,
            method:"get",
            height:500,
            singleSelect: true,
            onLoadSuccess:function(data){
                console.log("data",data)
            },
            loadFilter: function(data){
                data.total = data.Total
                data.rows = data.Rows
                return data
            },
            columns:[[
                {field:'id',title:'按钮Id',width:100},
                {field:'buttonName',title:'按钮名字',width:200},
                {field:'shell',title:'执行的shell指令',width:500,align:'right'}
            ]],toolbar:[{
                text:'增加',
                iconCls:'icon-add',
                handler:function () {
                    addButtonConfig()
                }},{
                text:'删除',
                iconCls:'icon-remove',
                handler:function () {
                    deleteButtonConfig()
                }},{
                text:'执行shell',
                iconCls:'icon-search',
                handler:function () {
                    sendShell()
                }}
                ]
        });
    }
    $("#combobox_ConfigButton").combobox({
        url:'/listSSHCombobox',
        valueField:'SSHId',
        method:"get",
        textField:'HostName',
        onSelect:function(record){
            $("#SSHId").val(record.SSHId)
        }
    })
// ,{
//         text:'执行',
//             iconCls:'icon-search',
//             handler:function () {
//             ScollPostionBase()
//         }}
    function ScollPostionBase() {
        var a = document.getElementById('scrol')
        a.scrollTop = a.scrollHeight
    }

    $('#addDialog').dialog({
        title: '添加',
        width: 400,
        height: 200,
        closed: true,
        cache: false,
        modal: true,
        toolbar:[{
            text:'保存',
            iconCls:'icon-save',
            handler:function () {
                addDialogSave()
            }},{
            text:'关闭',
            iconCls:'icon-close',
            handler:function () {
                $("#addDialog").dialog('close')
            }
         }]
    });


    /**
     * 添加 按钮配置
     */
    function addButtonConfig() {
        $("#addDialog").dialog('open')
    }

    /**
     * 执行 shell命令
     */
    function sendShell(){
        let datagrid = $('#tab').datagrid('getSelected');
        if (!datagrid){
            return
        }
        let SSHId = $("#SSHId").val();
        if (!SSHId) {
            alert('请选择一个SSH配置')
            return
        }
        $.messager.progress({
            title: '请稍等',
            msg: '执行中...',
        });

        $.ajax({
            url:"/executeShell",
            type:"get",
            data: {shell:datagrid.shell,SSHId:SSHId},
            success:function(data){
                // $("#tb").textbox("setValue",formatLog(data.Data.data))
                $("#code").append(formatLog(data.Data.data))
                ScollPostionBase()
                //关闭遮罩
                $.messager.progress('close');
            }

        })
    }



    /**
     * 删除 按钮配置
     */
    function deleteButtonConfig(){
        let datagrid = $('#tab').datagrid('getSelected');
        if (!datagrid){
            return
        }
        $.ajax({
            url:"/deleteById",
            type:"get",
            data: {id:datagrid.id},
            success:function(data){
                $("#tab").datagrid('reload');
            }

        })
    }
    /**
     *  添加按钮配置的二级对话框的保存按钮
     */
    function addDialogSave(){
        $('#addForm').form('submit', {
            url: "/saveAddDialog",
            onSubmit: function () {
                var isValid = $(this).form('validate');
                if (!isValid) {

                }
                fileData = $(this).serialize();
                return isValid;
            },
            success: function (result) {
                $("#tab").datagrid('reload');
                $("#addDialog").dialog('close');

            }
        });


    }
        let specialStr = {//特殊字符映射
            // " ":"&nbsp;",
            // "\t":"<br>",
            // "\n":"<br>",
            "[31m":"<font color='red'>",
            "[32m":"<font color='green'>",
            "[33m":"<font color='orange'>",
            "[34m":"<font color='#00008b'>",
            "[35m":"<font color='#006400'>",
            "[36m":"<font color='blue'>",
            "[37m":"<font color='#663399'>",
            "[39m":"</font>",
        };

        function formatLog(str){
            let tempStr = str;
            tempStr = tempStr.replace(new RegExp(" ",("gm")),"&nbsp;");
            tempStr = tempStr.replace(new RegExp("\t",("gm")),"<br>");
            tempStr = tempStr.replace(new RegExp("\n",("gm")),"<br>");
            $.each(specialStr,function(index,value){
                while (tempStr.indexOf(index) >= 0){
                    tempStr = tempStr.replace(index, value);
                }
                //tempStr = tempStr.replace(new RegExp(index,"gm"),value);
            });
            return tempStr;
        }
})
