//本服配置
let server = {
    id: "#serverId#",
    httpPort: "#serverHttpPort#",
    ip: "#serverIp#",
    name: "#serverName#",
    openDate: "#serverOpenDate#",
    version: "#serverVersion#",
    platform: "#serverPlatform#",
    partner: "#serverPartner#",
};

//充值服配置
let rechargeServer = {
    port: "#rechargeServerPort#",
    ip: "#rechargeServerIp#",
};

//账号服配置
let acctServer = {
    port: "#acctServerPort#",
    ip: "#acctServerIp#",
};


let mongo = {
    db: {
        usr: "#mongoUser#",
        pwd: "#mongoPwd#",
        host: "#mongoHost#",
        name: "SQZZ_" + server.id
    },
    chatDB: {
        usr: "#chatDBUser#",
        pwd: "#chatDBPwd#",
        host: "#chatDBHost#",
        name: "SQZZ_CHAT_" + server.id
    }
};

let mysql = {
    GameLog: {
        host: "#mysqlHost#",
        port: "#mysqlPort#",
        database: "#mysqlDataBase#",
        user: "#mysqlUser#",
        password: "#mysqlPassword#"
    }
};
let redis = {};


module.exports = {
    mongo:mongo,
    mysql:mysql,
    redis:redis,
    serverCfg: {
        "id": server.id,
        "name": server.name,
        "openDate": server.openDate,
        "version": server.version,
        "platform": server.platform,
        "partner": server.partner,
        "httpPort": server.httpPort,
        "secret": "9322C223A5D58",
        "syncPlayerInfoUrl": "http://" + acctServer.ip + ":" + acctServer.port + "/syncPlayerInfo",
        "updateRegisterCntUrl": "http://" + acctServer.ip + ":" + acctServer.port + "/updateRegisterCnt",
        "umengAppKey": "5b8eb6e48f4a9d758b000128",
        "umengSecret": "jeotyvg13jaefx94bwsbjg9ha5lsu581",
        "umengAppKeyIOS": "5c8a22df2036571232000b05",
        "umengSecretIOS": "me5egpsvsmaok6qjtmg6oo4ganoix12w",
        "umengUrl": "https://msgapi.umeng.com/api/send",
        "umengProduction": false,
        "rechargeUrl": "http://" + rechargeServer.ip + ":" + rechargeServer.port + "/makeOrder",
        "rechargeBackUrl": "http://" + rechargeServer.ip + ":" + rechargeServer.port + "/rechargeBack/checkAdd",
        "rechargeCallBackUrl": "http://" + server.ip + ":" + server.httpPort + "/rechargeCallBack"
    },

    "official": {
        "checkLoginUrl": "http://" + acctServer.ip + ":" + acctServer.port + "/checkLogin",
        "checkLoginMethod": "POST"
    },

    "smart": {
        "checkLoginUrl": "http://api.17173g.cn:28080/platform-login/user-login/checkLogin.do",
        "checkLoginMethod": "POST",
        "appid": "1efcc1ddd61afead4c699e66baea8fd8",
        "channellevel1": "MHT",
        "channellevel2": "MHT",
        "secretKey": "ef3c325ec6a534f215c1a3b1e42a8e8b"
    },

    "smartIOS": {
        "checkLoginUrl": "http://api.17173g.cn:28080/platform-login/user-login/checkLogin.do",
        "checkLoginMethod": "POST",
        "appid": "5c9680619fbe1f4dd0946ef8f0b06cd0",
        "channellevel1": "MHT",
        "channellevel2": "MHT",
        "secretKey": "c24c933808b590852033879e974b57f8"
    },


    "smartQuick": {
        "checkLoginUrl": "http://checkuser.sdk.quicksdk.net/v2/checkUserInfo",
        "checkLoginMethod": "POST",
        "ProductCode": "86900824234842295433832095360749",
        "ProductKey": "43942184",
        "Callback_Key": "52295590023856951212052491913768",
        "Md5_Key": "4urasmet00lbfg9qxk8zvl1sww3sigln"
    }
};