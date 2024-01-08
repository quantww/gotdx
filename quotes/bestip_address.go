package quotes

var (
	// StandardServerList 标准行情服务器列表
	StandardServerList = []Server{
		{Source: "通达信", Name: "深圳双线主站1", Host: "110.41.147.114", Port: 7709, CrossTime: 0},
		{Source: "通达信", Name: "深圳双线主站2", Host: "110.41.2.72", Port: 7709, CrossTime: 0},
		{Source: "通达信", Name: "深圳双线主站3", Host: "110.41.4.4", Port: 7709, CrossTime: 0},
		{Source: "通达信", Name: "深圳双线主站4", Host: "47.113.94.204", Port: 7709, CrossTime: 0},
		{Source: "通达信", Name: "深圳双线主站5", Host: "8.129.174.169", Port: 7709, CrossTime: 0},
		{Source: "通达信", Name: "深圳双线主站6", Host: "110.41.154.219", Port: 7709, CrossTime: 0},
		{Source: "通达信", Name: "上海双线主站1", Host: "124.70.176.52", Port: 7709, CrossTime: 0},
		{Source: "通达信", Name: "上海双线主站2", Host: "47.100.236.28", Port: 7709, CrossTime: 0},
		{Source: "通达信", Name: "上海双线主站3", Host: "123.60.186.45", Port: 7709, CrossTime: 0},
		{Source: "通达信", Name: "上海双线主站4", Host: "123.60.164.122", Port: 7709, CrossTime: 0},
		{Source: "通达信", Name: "上海双线主站5", Host: "47.116.105.28", Port: 7709, CrossTime: 0},
		{Source: "通达信", Name: "上海双线主站6", Host: "124.70.199.56", Port: 7709, CrossTime: 0},
		{Source: "通达信", Name: "北京双线主站1", Host: "121.36.54.217", Port: 7709, CrossTime: 0},
		{Source: "通达信", Name: "北京双线主站2", Host: "121.36.81.195", Port: 7709, CrossTime: 0},
		{Source: "通达信", Name: "北京双线主站3", Host: "123.249.15.60", Port: 7709, CrossTime: 0},
		{Source: "通达信", Name: "广州双线主站1", Host: "124.71.85.110", Port: 7709, CrossTime: 0},
		{Source: "通达信", Name: "广州双线主站2", Host: "139.9.51.18", Port: 7709, CrossTime: 0},
		{Source: "通达信", Name: "广州双线主站3", Host: "139.159.239.163", Port: 7709, CrossTime: 0},
		{Source: "通达信", Name: "上海双线主站7", Host: "106.14.201.131", Port: 7709, CrossTime: 0},
		{Source: "通达信", Name: "上海双线主站8", Host: "106.14.190.242", Port: 7709, CrossTime: 0},
		{Source: "通达信", Name: "上海双线主站9", Host: "121.36.225.169", Port: 7709, CrossTime: 0},
		{Source: "通达信", Name: "上海双线主站10", Host: "123.60.70.228", Port: 7709, CrossTime: 0},
		{Source: "通达信", Name: "上海双线主站11", Host: "123.60.73.44", Port: 7709, CrossTime: 0},
		{Source: "通达信", Name: "上海双线主站12", Host: "124.70.133.119", Port: 7709, CrossTime: 0},
		{Source: "通达信", Name: "上海双线主站13", Host: "124.71.187.72", Port: 7709, CrossTime: 0},
		{Source: "通达信", Name: "上海双线主站14", Host: "124.71.187.122", Port: 7709, CrossTime: 0},
		{Source: "通达信", Name: "武汉电信主站1", Host: "119.97.185.59", Port: 7709, CrossTime: 0},
		{Source: "通达信", Name: "深圳双线主站7", Host: "47.107.64.168", Port: 7709, CrossTime: 0},
		{Source: "通达信", Name: "北京双线主站4", Host: "124.70.75.113", Port: 7709, CrossTime: 0},
		{Source: "通达信", Name: "广州双线主站4", Host: "124.71.9.153", Port: 7709, CrossTime: 0},
		{Source: "通达信", Name: "上海双线主站15", Host: "123.60.84.66", Port: 7709, CrossTime: 0},
		{Source: "通达信", Name: "深圳双线主站8", Host: "47.107.228.47", Port: 7719, CrossTime: 0},
		{Source: "通达信", Name: "北京双线主站5", Host: "120.46.186.223", Port: 7709, CrossTime: 0},
		{Source: "通达信", Name: "北京双线主站6", Host: "124.70.22.210", Port: 7709, CrossTime: 0},
		{Source: "通达信", Name: "北京双线主站7", Host: "139.9.133.247", Port: 7709, CrossTime: 0},
		{Source: "通达信", Name: "广州双线主站5", Host: "116.205.163.254", Port: 7709, CrossTime: 0},
		{Source: "通达信", Name: "广州双线主站6", Host: "116.205.171.132", Port: 7709, CrossTime: 0},
		{Source: "通达信", Name: "广州双线主站7", Host: "116.205.183.150", Port: 7709, CrossTime: 0},
		{Source: "中信证券", Name: "上海电信主站Z1", Host: "180.153.18.170", Port: 7709, CrossTime: 0},
		{Source: "中信证券", Name: "上海电信主站Z2", Host: "180.153.18.171", Port: 7709, CrossTime: 0},
		{Source: "中信证券", Name: "北京联通主站Z1", Host: "202.108.253.130", Port: 7709, CrossTime: 0},
		{Source: "中信证券", Name: "北京联通主站Z2", Host: "202.108.253.131", Port: 7709, CrossTime: 0},
		{Source: "中信证券", Name: "杭州电信主站J1", Host: "60.191.117.167", Port: 7709, CrossTime: 0},
		{Source: "中信证券", Name: "杭州电信主站J2", Host: "115.238.56.198", Port: 7709, CrossTime: 0},
		{Source: "中信证券", Name: "杭州电信主站J3", Host: "218.75.126.9", Port: 7709, CrossTime: 0},
		{Source: "中信证券", Name: "杭州电信主站J4", Host: "115.238.90.165", Port: 7709, CrossTime: 0},
		{Source: "中信证券", Name: "杭州联通主站J1", Host: "124.160.88.183", Port: 7709, CrossTime: 0},
		{Source: "中信证券", Name: "杭州联通主站J2", Host: "60.12.136.250", Port: 7709, CrossTime: 0},
		{Source: "中信证券", Name: "杭州华数主站J1", Host: "218.108.98.244", Port: 7709, CrossTime: 0},
		{Source: "中信证券", Name: "杭州华数主站J2", Host: "218.108.47.69", Port: 7709, CrossTime: 0},
		{Source: "中信证券", Name: "济南联通主站W1", Host: "27.221.115.131", Port: 7709, CrossTime: 0},
		{Source: "中信证券", Name: "青岛电信主站W1", Host: "58.56.180.60", Port: 7709, CrossTime: 0},
		{Source: "中信证券", Name: "深圳电信主站Z1", Host: "14.17.75.71", Port: 7709, CrossTime: 0},
		{Source: "中信证券", Name: "云行情上海电信Z1", Host: "114.80.63.12", Port: 7709, CrossTime: 0},
		{Source: "中信证券", Name: "云行情上海电信Z2", Host: "114.80.63.35", Port: 7709, CrossTime: 0},
		{Source: "中信证券", Name: "上海电信主站Z3", Host: "180.153.39.51", Port: 7709, CrossTime: 0},
		{Source: "中信证券", Name: "云行情北京联通Z1", Host: "123.125.108.23", Port: 7709, CrossTime: 0},
		{Source: "中信证券", Name: "云行情北京联通Z2", Host: "123.125.108.24", Port: 7709, CrossTime: 0},
		{Source: "中信证券", Name: "云行情广州电信Z1", Host: "121.201.83.106", Port: 7709, CrossTime: 0},
		{Source: "中信证券", Name: "云行情成都电信Z1", Host: "218.6.170.55", Port: 7709, CrossTime: 0},
		{Source: "华泰证券", Name: "华泰证券(智选一)", Host: "tdxhq.htzq.com.cn", Port: 7709, CrossTime: 0},
		{Source: "华泰证券", Name: "华泰证券(智选二)", Host: "tdxhq.htsc.com", Port: 7709, CrossTime: 0},
		{Source: "华泰证券", Name: "华泰证券(南京电信一)", Host: "180.101.48.170", Port: 7709, CrossTime: 0},
		{Source: "华泰证券", Name: "华泰证券(南京电信二)", Host: "180.101.48.171", Port: 7709, CrossTime: 0},
		{Source: "华泰证券", Name: "华泰证券(南京移动一)", Host: "120.195.71.155", Port: 7709, CrossTime: 0},
		{Source: "华泰证券", Name: "华泰证券(南京移动二)", Host: "120.195.71.156", Port: 7709, CrossTime: 0},
		{Source: "华泰证券", Name: "华泰证券(南京联通一)", Host: "122.96.107.242", Port: 7709, CrossTime: 0},
		{Source: "华泰证券", Name: "华泰证券(南京联通二)", Host: "122.96.107.243", Port: 7709, CrossTime: 0},
		{Source: "华泰证券", Name: "华泰证券(亚马逊一)", Host: "52.83.39.241", Port: 7709, CrossTime: 0},
		{Source: "华泰证券", Name: "华泰证券(亚马逊二)", Host: "52.83.199.101", Port: 7709, CrossTime: 0},
		{Source: "华泰证券", Name: "华泰证券(华南阿里云一)", Host: "8.135.57.58", Port: 7709, CrossTime: 0},
		{Source: "华泰证券", Name: "华泰证券(华南阿里云二)", Host: "8.135.62.177", Port: 7709, CrossTime: 0},
		{Source: "华泰证券", Name: "华泰证券(华东华为云一)", Host: "124.70.183.173", Port: 7709, CrossTime: 0},
		{Source: "华泰证券", Name: "华泰证券(华东华为云二)", Host: "124.71.163.106", Port: 7709, CrossTime: 0},
		{Source: "国泰君安", Name: "国泰君安深圳电信组站", Host: "sztdx.gtjas.com", Port: 7709, CrossTime: 0},
		{Source: "国泰君安", Name: "国泰君安江苏电信组站", Host: "jstdx.gtjas.com", Port: 7709, CrossTime: 0},
		{Source: "国泰君安", Name: "国泰君安上海电信组站", Host: "shtdx.gtjas.com", Port: 7709, CrossTime: 0},
		{Source: "国泰君安", Name: "国泰君安北京网通组站", Host: "bjwttdx.gtjas.com", Port: 7709, CrossTime: 0},
		{Source: "国泰君安", Name: "国泰君安北方网通组站", Host: "bfwttdx.gtjas.com", Port: 7709, CrossTime: 0},
		{Source: "国泰君安", Name: "国泰君安河北网通组站", Host: "hbwttdx.gtjas.com", Port: 7709, CrossTime: 0},
		{Source: "国泰君安", Name: "国泰君安成都电信组站", Host: "cdtdx.gtjas.com", Port: 7709, CrossTime: 0},
		{Source: "国泰君安", Name: "郑州网通行情一", Host: "182.118.47.141", Port: 7709, CrossTime: 0},
		{Source: "国泰君安", Name: "郑州网通行情二", Host: "182.118.47.168", Port: 7709, CrossTime: 0},
		{Source: "国泰君安", Name: "郑州网通行情三", Host: "182.118.47.169", Port: 7709, CrossTime: 0},
		{Source: "国泰君安", Name: "武汉电信行情一", Host: "119.97.164.184", Port: 7709, CrossTime: 0},
		{Source: "国泰君安", Name: "武汉电信行情二", Host: "119.97.164.189", Port: 7709, CrossTime: 0},
		{Source: "国泰君安", Name: "武汉电信行情三", Host: "116.211.121.102", Port: 7709, CrossTime: 0},
		{Source: "国泰君安", Name: "武汉电信行情四", Host: "116.211.121.108", Port: 7709, CrossTime: 0},
		{Source: "国泰君安", Name: "武汉电信行情五", Host: "116.211.121.31", Port: 7709, CrossTime: 0},
		{Source: "国泰君安", Name: "新疆电信云行情一", Host: "202.100.166.117", Port: 7709, CrossTime: 0},
		{Source: "国泰君安", Name: "新疆电信云行情二", Host: "202.100.166.118", Port: 7709, CrossTime: 0},
		{Source: "国泰君安", Name: "上海电信行情八", Host: "222.73.139.166", Port: 7709, CrossTime: 0},
		{Source: "国泰君安", Name: "上海电信行情九", Host: "222.73.139.167", Port: 7709, CrossTime: 0},
		{Source: "国泰君安", Name: "上海电信行情十", Host: "222.73.139.168", Port: 7709, CrossTime: 0},
		{Source: "国泰君安", Name: "上海BGP行情一", Host: "103.251.85.90", Port: 7709, CrossTime: 0},
		{Source: "国泰君安", Name: "北京联通行情一", Host: "123.125.108.213", Port: 7709, CrossTime: 0},
		{Source: "国泰君安", Name: "北京联通行情二", Host: "123.125.108.214", Port: 7709, CrossTime: 0},
		{Source: "国泰君安", Name: "上海电信行情六", Host: "222.73.139.151", Port: 7709, CrossTime: 0},
		{Source: "国泰君安", Name: "上海电信行情七", Host: "222.73.139.152", Port: 7709, CrossTime: 0},
		{Source: "国泰君安", Name: "成都BGP行情一", Host: "148.70.110.41", Port: 7709, CrossTime: 0},
		{Source: "国泰君安", Name: "成都BGP行情二", Host: "148.70.93.117", Port: 7709, CrossTime: 0},
		{Source: "国泰君安", Name: "成都BGP行情三", Host: "148.70.31.16", Port: 7709, CrossTime: 0},
		{Source: "国泰君安", Name: "成都BGP行情四", Host: "148.70.111.63", Port: 7709, CrossTime: 0},
		{Source: "国泰君安", Name: "广州BGP行情一", Host: "139.159.143.228", Port: 7709, CrossTime: 0},
		{Source: "国泰君安", Name: "广州BGP行情二", Host: "139.159.183.76", Port: 7709, CrossTime: 0},
		{Source: "国泰君安", Name: "广州BGP行情三", Host: "139.159.193.118", Port: 7709, CrossTime: 0},
		{Source: "国泰君安", Name: "广州BGP行情四", Host: "139.159.195.177", Port: 7709, CrossTime: 0},
		{Source: "国泰君安", Name: "广州BGP行情五", Host: "139.159.202.253", Port: 7709, CrossTime: 0},
		{Source: "国泰君安", Name: "广州BGP行情六", Host: "139.159.214.78", Port: 7709, CrossTime: 0},
		{Source: "国泰君安", Name: "广州BGP行情七", Host: "139.9.38.206", Port: 7709, CrossTime: 0},
		{Source: "国泰君安", Name: "广州BGP行情八", Host: "139.9.43.104", Port: 7709, CrossTime: 0},
		{Source: "国泰君安", Name: "广州BGP行情九", Host: "139.9.43.31", Port: 7709, CrossTime: 0},
		{Source: "国泰君安", Name: "广州BGP行情十", Host: "139.9.50.246", Port: 7709, CrossTime: 0},
		{Source: "国泰君安", Name: "广州BGP行情十一", Host: "139.9.52.158", Port: 7709, CrossTime: 0},
		{Source: "国泰君安", Name: "广州BGP行情十二", Host: "139.9.90.169", Port: 7709, CrossTime: 0},
		{Source: "国泰君安", Name: "上海电信行情十一", Host: "101.226.180.73", Port: 7709, CrossTime: 0},
		{Source: "国泰君安", Name: "上海电信行情十二", Host: "101.226.180.74", Port: 7709, CrossTime: 0},
		{Source: "国泰君安", Name: "上海BGP行情六", Host: "103.251.85.200", Port: 7709, CrossTime: 0},
		{Source: "国泰君安", Name: "上海BGP行情七", Host: "103.251.85.201", Port: 7709, CrossTime: 0},
		{Source: "国泰君安", Name: "南京电信行情一", Host: "103.221.142.65", Port: 7709, CrossTime: 0},
		{Source: "国泰君安", Name: "南京电信行情二", Host: "103.221.142.66", Port: 7709, CrossTime: 0},
		{Source: "国泰君安", Name: "南京电信行情三", Host: "103.221.142.67", Port: 7709, CrossTime: 0},
		{Source: "国泰君安", Name: "南京电信行情四", Host: "103.221.142.68", Port: 7709, CrossTime: 0},
		{Source: "国泰君安", Name: "南京电信行情五", Host: "103.221.142.69", Port: 7709, CrossTime: 0},
		{Source: "国泰君安", Name: "南京电信行情六", Host: "103.221.142.70", Port: 7709, CrossTime: 0},
		{Source: "国泰君安", Name: "南京电信行情七", Host: "103.221.142.71", Port: 7709, CrossTime: 0},
		{Source: "国泰君安", Name: "南京电信行情八", Host: "103.221.142.72", Port: 7709, CrossTime: 0},
		{Source: "国泰君安", Name: "西安电信行情一", Host: "117.34.114.13", Port: 7709, CrossTime: 0},
		{Source: "国泰君安", Name: "西安电信行情二", Host: "117.34.114.14", Port: 7709, CrossTime: 0},
		{Source: "国泰君安", Name: "西安电信行情三", Host: "117.34.114.15", Port: 7709, CrossTime: 0},
		{Source: "国泰君安", Name: "西安电信行情四", Host: "117.34.114.16", Port: 7709, CrossTime: 0},
		{Source: "国泰君安", Name: "西安电信行情五", Host: "117.34.114.17", Port: 7709, CrossTime: 0},
		{Source: "国泰君安", Name: "西安电信行情六", Host: "117.34.114.18", Port: 7709, CrossTime: 0},
		{Source: "国泰君安", Name: "西安电信行情七", Host: "117.34.114.20", Port: 7709, CrossTime: 0},
		{Source: "国泰君安", Name: "西安电信行情八", Host: "117.34.114.27", Port: 7709, CrossTime: 0},
		{Source: "国泰君安", Name: "西安电信行情九", Host: "117.34.114.30", Port: 7709, CrossTime: 0},
		{Source: "国泰君安", Name: "上海BGP行情八", Host: "103.251.85.202", Port: 7709, CrossTime: 0},
		{Source: "国泰君安", Name: "东莞电信行情一", Host: "183.60.224.142", Port: 7709, CrossTime: 0},
		{Source: "国泰君安", Name: "东莞电信行情二", Host: "183.60.224.143", Port: 7709, CrossTime: 0},
		{Source: "国泰君安", Name: "东莞电信行情三", Host: "183.60.224.144", Port: 7709, CrossTime: 0},
		{Source: "国泰君安", Name: "东莞电信行情四", Host: "183.60.224.145", Port: 7709, CrossTime: 0},
		{Source: "国泰君安", Name: "东莞电信行情五", Host: "183.60.224.146", Port: 7709, CrossTime: 0},
		{Source: "国泰君安", Name: "东莞电信行情六", Host: "183.60.224.147", Port: 7709, CrossTime: 0},
		{Source: "国泰君安", Name: "东莞电信行情七", Host: "183.60.224.148", Port: 7709, CrossTime: 0},
	}
	// ExtensionServerList 扩展行情服务器列表
	ExtensionServerList = []Server{
		{Source: "通达信", Name: "扩展市场深圳双线1", Host: "112.74.214.43", Port: 7727, CrossTime: 0},
		{Source: "通达信", Name: "扩展市场深圳双线2", Host: "120.25.218.6", Port: 7727, CrossTime: 0},
		{Source: "通达信", Name: "扩展市场深圳双线3", Host: "47.107.75.159", Port: 7727, CrossTime: 0},
		{Source: "通达信", Name: "扩展市场深圳双线4", Host: "47.106.204.218", Port: 7727, CrossTime: 0},
		{Source: "通达信", Name: "扩展市场深圳双线5", Host: "47.106.209.131", Port: 7727, CrossTime: 0},
		{Source: "通达信", Name: "扩展市场武汉主站1", Host: "119.97.185.5", Port: 7727, CrossTime: 0},
		{Source: "通达信", Name: "扩展市场深圳双线6", Host: "47.115.94.72", Port: 7727, CrossTime: 0},
		{Source: "通达信", Name: "扩展市场上海双线1", Host: "106.14.95.149", Port: 7727, CrossTime: 0},
		{Source: "通达信", Name: "扩展市场上海双线2", Host: "47.102.108.214", Port: 7727, CrossTime: 0},
		{Source: "通达信", Name: "扩展市场上海双线3", Host: "47.103.86.229", Port: 7727, CrossTime: 0},
		{Source: "通达信", Name: "扩展市场上海双线4", Host: "47.103.88.146", Port: 7727, CrossTime: 0},
		{Source: "通达信", Name: "扩展市场广州双线1", Host: "116.205.143.214", Port: 7727, CrossTime: 0},
		{Source: "通达信", Name: "扩展市场广州双线2", Host: "124.71.223.19", Port: 7727, CrossTime: 0},
		{Source: "中信证券", Name: "上海电信主站Z1", Host: "180.153.18.176", Port: 7721, CrossTime: 0},
		{Source: "中信证券", Name: "北京联通主站Z1", Host: "202.108.253.154", Port: 7721, CrossTime: 0},
		{Source: "中信证券", Name: "杭州电信主站J1", Host: "115.238.56.196", Port: 7721, CrossTime: 0},
		{Source: "中信证券", Name: "杭州电信主站J2", Host: "115.238.90.170", Port: 7721, CrossTime: 0},
		{Source: "中信证券", Name: "杭州联通主站J1", Host: "60.12.136.251", Port: 7721, CrossTime: 0},
		{Source: "中信证券", Name: "杭州华数主站J1", Host: "218.108.98.244", Port: 7721, CrossTime: 0},
		{Source: "中信证券", Name: "济南联通主站W1", Host: "27.221.115.133", Port: 7721, CrossTime: 0},
		{Source: "中信证券", Name: "青岛电信主站W1", Host: "58.56.180.60", Port: 7721, CrossTime: 0},
		{Source: "中信证券", Name: "深圳电信主站Z1", Host: "14.17.75.71", Port: 7721, CrossTime: 0},
		{Source: "中信证券", Name: "广州云电信主站Z1", Host: "121.201.83.104", Port: 7721, CrossTime: 0},
		{Source: "华泰证券", Name: "华泰证券(智选一)", Host: "tdxkzhq.htzq.com.cn", Port: 7721, CrossTime: 0},
		{Source: "华泰证券", Name: "华泰证券(智选二)", Host: "tdxkzhq.htsc.com", Port: 7721, CrossTime: 0},
		{Source: "华泰证券", Name: "华泰证券(南京电信一)", Host: "180.101.48.170", Port: 7721, CrossTime: 0},
		{Source: "华泰证券", Name: "华泰证券(南京电信二)", Host: "180.101.48.171", Port: 7721, CrossTime: 0},
		{Source: "华泰证券", Name: "华泰证券(南京移动一)", Host: "120.195.71.155", Port: 7721, CrossTime: 0},
		{Source: "华泰证券", Name: "华泰证券(南京移动二)", Host: "120.195.71.156", Port: 7721, CrossTime: 0},
		{Source: "华泰证券", Name: "华泰证券(南京联通一)", Host: "122.96.107.242", Port: 7721, CrossTime: 0},
		{Source: "华泰证券", Name: "华泰证券(南京联通二)", Host: "122.96.107.243", Port: 7721, CrossTime: 0},
		{Source: "华泰证券", Name: "华泰证券(亚马逊一)", Host: "52.83.39.241", Port: 7721, CrossTime: 0},
		{Source: "华泰证券", Name: "华泰证券(亚马逊二)", Host: "52.83.199.101", Port: 7721, CrossTime: 0},
		{Source: "华泰证券", Name: "华泰证券(华南阿里云一)", Host: "8.135.57.58", Port: 7721, CrossTime: 0},
		{Source: "华泰证券", Name: "华泰证券(华南阿里云二)", Host: "8.135.62.177", Port: 7721, CrossTime: 0},
		{Source: "华泰证券", Name: "华泰证券(华东华为云一)", Host: "124.70.183.173", Port: 7721, CrossTime: 0},
		{Source: "华泰证券", Name: "华泰证券(华东华为云二)", Host: "124.71.163.106", Port: 7721, CrossTime: 0},
		{Source: "国泰君安", Name: "扩展行情主站1", Host: "103.221.142.80", Port: 7721, CrossTime: 0},
		{Source: "国泰君安", Name: "扩展行情主站2", Host: "114.118.82.205", Port: 7721, CrossTime: 0},
		{Source: "国泰君安", Name: "扩展行情主站3", Host: "117.34.114.31", Port: 7721, CrossTime: 0},
		{Source: "国泰君安", Name: "扩展行情主站4", Host: "139.9.52.158", Port: 7721, CrossTime: 0},
		{Source: "国泰君安", Name: "扩展行情主站5", Host: "103.251.85.204", Port: 7721, CrossTime: 0},
		{Source: "国泰君安", Name: "扩展行情主站6", Host: "114.118.82.204", Port: 7721, CrossTime: 0},
		{Source: "国泰君安", Name: "扩展行情主站7", Host: "103.221.142.73", Port: 7721, CrossTime: 0},
	}
)
