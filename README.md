#### 目录介绍
    application 应用层 该层用作域和界面层之间的通道。将请求从接口层发送到域层，由域层处理请求并返回响应
     
    domain 领域层
        aggregates 聚合
		entity 实体
		po 与数据库交互
		valueobj 值对象
		vo  View Object
	
	infra Infrastructure基础设施层
	    config 项目相关配置
	    util 主要放数据库、缓存、文件、网关等第三方类库
	
	interfaces 该层包含与其他系统交互的所有内容，例如Web服务
	    assembler 组装 实现DTO 与领域对象之间的相互转换、映射
	    dto 数据传输的载体，我们可以通过DTO把内部的领域对象与外界隔离
	    
	server 
	    main.go 项目运行脚本
	    
	