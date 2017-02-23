$(function(){

	// tab切换
	function tab(aa,bb){
		for(var i=0;i<aa.length;i++){
			aa.eq(i).click(function(){
				var index=$(this).index()
				$(this).addClass("active").siblings().removeClass("active")
				bb.eq(index).addClass("show").removeClass("hide").siblings().removeClass("show").addClass("hide")
			})
		}
	}
	// tab功能
	tab($(".ul li"),$(".tab>div"))

// 刘龙勤大区ajax复选框
	function tim(aa){
		aa.html("")
		
		$.ajax({
			type:"post",
			url:"../query_area_servers.php",
			dataType:"json",
			success:function(e){
				for(var key in e){
					$('.legend').html("");
					$('.legend').append(" "+key)
					var tim=e[key]
					for(var i=0; i<tim.length;i++){
						 var id=tim[i].id
						 lable= '<input type="checkbox" name="servs"  value="'+id+'"   class="ui-checkboxradio ui-helper-hidden-accessible"/><label for="service_notice_serv_1" class="ui-button ui-widget ui-controlgroup-item ui-checkboxradio-label ui-corner-left"><span class="ui-checkboxradio-icon ui-corner-all ui-icon ui-icon-background ui-icon-blank"></span><span class="ui-checkboxradio-icon-space"></span>'+tim[i].name+'</label> '
						aa.append(lable)
					}
				}
			}
		})
	}
	tim($(".public"))

	//单选框
	function tim1(aa){
		aa.html("")
		$.ajax({
			type:"post",
			url:"../query_area_servers.php",
			dataType:"json",
			success:function(e){
				for(var key in e){
					$('.legend').html("");
					$('.legend').append(" "+key)
					var tim=e[key]
					for(var i=0; i<tim.length;i++){
						 var id=tim[i].id
						 lable= '<input type="radio" name="servs" value="'+id+'"  class="ui-checkboxradio ui-helper-hidden-accessible"/><label for="service_notice_serv_1" class="ui-button ui-widget ui-controlgroup-item ui-checkboxradio-label ui-corner-left"><span class="ui-checkboxradio-icon ui-corner-all ui-icon ui-icon-background ui-icon-blank"></span><span class="ui-checkboxradio-icon-space"></span>'+tim[i].name+'</label> '
						aa.append(lable)
					}
				}
			}
		})
	}
	tim1($(".public1"))
	
	
	
	//订单页面选框
	
		$.ajax({
			type:"post",
			url:"../query_area_servers.php",
			dataType:"json",
			success:function(e){
				for(var key in e){
					var tim=e[key]
					for(var i=0; i<tim.length;i++){
						$(".order_input").val(tim[0].name)
						$(".pay_input").val(tim[0].name)
						 var id=tim[i].id
						 lable= '<li id="'+id+'"  class="option" style="cursor:pointer;list-style: none;height: 30px;width: 173px;margin-left: -35px;" >'+tim[i].name+'</li>'
						$("#sever_types").append(lable)
						$("#pay_types").append(lable)
					}
				} 
				
			}
		})
		//订单统计
		
		$(".order_input").click(function(){
			$("#sever_types").removeClass("hide").addClass("show")
		})
		$(document).on("click",".option",function(){
			var val=$(".order_id").val()
			$(".order_input").val($(this).html())
			$(this).parent().removeClass("show").addClass("hide")
			var id=$(this).attr('id')
			$.ajax({
				type:"post",
				url:"../order_inquiry.php",
				dataType:"json",
				data:{id:id,val:val},
				success:function(e){
					
				var data = [];
				for(var i=0;i<e.length;i++)
				{
					data[i] =[e[i].Game,e[i].Paytime,e[i].Pfid,e[i].Pfname,e[i].Orderid,e[i].Roleid,e[i].Accountid,e[i].Payment];
				}
				var cs = new table({
				 "tableId":"gs_table",    //必须 表格id
				"headers":["游戏名","支付时间","渠道id","渠道名","订单id","角色id","账户名","支付金额"],   //必须 thead表头
				"data":data,         //必须 tbody 数据展示
				"displayNum": 20,    //必须   默认 10  每页显示行数
				"groupDataNum":1,     //可选    默认 10  组数
				"display_tfoot":true, // true/false  是否显示tfoot --- 默认false
				"bindContentTr":function(){ //可选 给tbody 每行绑定事件回调
					this.tableObj.find("tbody").on("click",'tr',function(e){
					return false;
					var tr_index = $(this).data("tr_index");        // tr行号  从0开始
					 var data_index = $(this).data("data_index");   //数据行号  从0开始
					})
				},
				sort:true,    // 点击表头是否排序 true/false  --- 默认false
				search:true   // 默认为false 没有搜索
						 });
					}
				
			})
			
		})
		//付费统计
		//$(".pay_input").click(function(){
			//$("#pay_types").removeClass("hide").addClass("show")
		//})
		//$(document).on("click",".option",function(){
			//$(".pay_input").val($(this).html())
			//$(this).parent().removeClass("show").addClass("hide")
			//var id=$(this).attr('id')
			//$.ajax({
				//type:"post",
				//url:"../pay_statistics.php",
				//dataType:"json",
				//data:{id:id},
				//success:function(e){
					
				//var data = [];
				//for(var i=0;i<e.length;i++)
				//{
					//data[i] =[e[i].Sortid,e[i].Game,e[i].Pfid,e[i].Pfname,e[i].Orderid,e[i].Accountid,e[i].Payment,e[i].Paytime];
				//}
				//var cs = new table({
				 //"tableId":"fs_table",    //必须 表格id
				//"headers":["id","游戏名","渠道id","渠道名","订单id","账户名","支付金额","支付时间"],   //必须 thead表头
				//"data":data,         //必须 tbody 数据展示
				//"displayNum": 20,    //必须   默认 10  每页显示行数
				//"groupDataNum":1,     //可选    默认 10  组数
				//"display_tfoot":true, // true/false  是否显示tfoot --- 默认false
				//"bindContentTr":function(){ //可选 给tbody 每行绑定事件回调
					//this.tableObj.find("tbody").on("click",'tr',function(e){
					//return false;
					//var tr_index = $(this).data("tr_index");        // tr行号  从0开始
					// var data_index = $(this).data("data_index");   //数据行号  从0开始
					//})
				//},
				//sort:true,    // 点击表头是否排序 true/false  --- 默认false
				//search:true   // 默认为false 没有搜索
						// });
					//}
				
			//})
			
		//})
	
//	历史记录查询

//历史查询
	function htab(aa){
		for(var i=0;i<aa.length;i++){
			aa.eq(i).click(function(){
				$(this).addClass("bg").siblings().removeClass("bg")
				$(this).find("input[type='button']").addClass("bg")
				$(this).siblings().find("input[type='button']").removeClass("bg")
			})
		}
	}
	htab($(".nav_ul li"))

//	循环点击着个出现
	for(var i=0;i<$(".nav_ul li input[type='button']").length;i++){

		$(".nav_ul li input[type='button']").eq(i).click(function(){
			var val=$(this).val()
			var parent=$(this).parent()
			$.ajax({
			type:"post",
			url:"../history.php",
			dataType:"json",
			data:parent.serialize(),
			success:function(e){
				var arr=[]
				 var data = [];
				if(val=="历史系统公告查询"){
					for(var i=0;i<e.length;i++){
						data[i] =[e[i].Title,e[i].Content,e[i].Color,e[i].Time];

					}
					arr=["公告标题","公告内容","公告颜色","时间"]
				}
				if(val=="历史区服公告查询"){
					for(var i=0;i<e.length;i++){
						data[i] =[e[i].Title,e[i].Content,e[i].ContentColor,e[i].Time];
					}
					arr=["公告标题","公告内容","公告颜色","时间"]
				}
				if(val=="历史滚动公告查询"){
					for(var i=0;i<e.length;i++){
						data[i] =[e[i].SendType,e[i].Content,e[i].Timestr,e[i].Time];
					}
					arr=["发送类型","公告内容","时间参数","时间"]
				}
				if(val=="历史发送邮件查询"){
					for(var i=0;i<e.length;i++){
						data[i] =[e[i].Title,e[i].Sender,e[i].Content,e[i].Recvers,e[i].Stritemids,e[i].Stritemsks,e[i].Lowlevel,e[i].Highlevel,e[i].Time0,e[i].Time1,e[i].SendType,e[i].Time];
					}
					arr=["标题","发件人","邮件内容","收件人","道具","堆叠","最低等级","最高等级","最早注册","最晚注册","类型","时间"]
				}
				if(val=="登录活动状态查询"){
					for(var i=0;i<e.length;i++){
						data[i] =[e[i].Name,e[i].OpenTime,e[i].CloseTime,e[i].Status,e[i].Time];
					}
					arr=["活动名称","开始时间","结束时间","活动状态","时间"]
				}
				if(val=="累积充值状态查询"){
					for(var i=0;i<e.length;i++){
						data[i] =[e[i].Name,e[i].OpenTime,e[i].CloseTime,e[i].Status,e[i].Time];
					}
					arr=["活动名称","开始时间","结束时间","活动状态","时间"]
				}
				if(val=="打折商店状态查询"){
					for(var i=0;i<e.length;i++){
						data[i] =[e[i].Name,e[i].OpenTime,e[i].CloseTime,e[i].Status,e[i].Time];
					}
					arr=["活动名称","开始时间","结束时间","活动状态","时间"]
				}
				if(val=="单笔充值状态查询"){
					for(var i=0;i<e.length;i++){
						data[i] =[e[i].Name,e[i].OpenTime,e[i].CloseTime,e[i].Status,e[i].Time];
					}
					arr=["活动名称","开始时间","结束时间","活动状态","时间"]
				}
				if(val=="热点伙伴状态查询"){
					for(var i=0;i<e.length;i++){
						data[i] =[e[i].Name,e[i].OpenTime,e[i].CloseTime,e[i].Status,e[i].Time];
					}
					arr=["活动名称","开始时间","结束时间","活动状态","时间"]
				}
				if(val=="顶级招募状态查询"){
					for(var i=0;i<e.length;i++){
						data[i] =[e[i].Name,e[i].OpenTime,e[i].CloseTime,e[i].Status,e[i].Time];
					}
					arr=["活动名称","开始时间","结束时间","活动状态","时间"]
				}
				if(val=="小额礼包状态查询"){
					for(var i=0;i<e.length;i++){
						data[i] =[e[i].Name,e[i].OpenTime,e[i].CloseTime,e[i].Status,e[i].Time];
					}
					arr=["活动名称","开始时间","结束时间","活动状态","时间"]
				}
				 var cs = new table({
				        "tableId":"ss_table",    //必须 表格id
				        "headers":arr,   //必须 thead表头
				        "data":data,         //必须 tbody 数据展示
				        "displayNum": 20,    //必须   默认 10  每页显示行数
				        "groupDataNum":1,     //可选    默认 10  组数
				        "display_tfoot":true, // true/false  是否显示tfoot --- 默认false
				        "bindContentTr":function(){ //可选 给tbody 每行绑定事件回调
				            this.tableObj.find("tbody").on("click",'tr',function(e){
				                return false;
				                var tr_index = $(this).data("tr_index");        // tr行号  从0开始
				                var data_index = $(this).data("data_index");   //数据行号  从0开始
				            })
				        },
				        sort:true,    // 点击表头是否排序 true/false  --- 默认false
				        search:true   // 默认为false 没有搜索
				    });
				},
				error:function(e){
				alert("执行错误")
				}
			})
		})
	}
	// 提交表单
	function pushback(a,urls){
			//alert(a.serialize())
			$.ajax({
			type:"post",
			url:urls,
			dataType:"json",
			data:a.serialize(),
			success:function(e){
				alert("执行成功")
			},
			error:function(e){
				alert("执行错误")
			}
		})
	}
	// 区域公告模快ajax
	$("#btn").click(function(){
			pushback($("#server"),"../change_server_notice.php")
		})
	// 系统公告模块ajax
	$("#bot").click(	function(){
			pushback($("#system"),"../change_system_notice.php")
		})
	// 滚动公告ajax
	$("#bnt").click(function(){
			pushback($("#roll"),"../chang_roll_notice.php")
		})
	// 邮件发送ajax
	$("#but").click(function(){
			pushback($("#mail"),"../insert_mail.php")
		})
	// 角色操作ajax
	$("#btm").click(function(){
		pushback($("#player"),"../contro_player.php")
	})
	// 登录活动ajax
	$("#sub").click(function(){
		pushback($("#activity"),"../login_activity.php")
	})
	//累计充值ajax
	$("#btn2").click(function(){
		pushback($("#accumulate_recharge"),"../accumulate_recharge.php")
	})
	//打折商店ajax
	$("#btn3").click(function(){
		pushback($("#discount_store"),"../discount_store.php")
	})
	//单笔充值ajax
	$("#btn4").click(function(){
		pushback($("#single_recharge"),"../single_recharge.php")
	})
	//热点伙伴ajax
	$("#btn5").click(function(){
		pushback($("#hotspot_partner"),"../hotspot_partner.php")
	})
	//顶级招募伙伴ajax
	$("#btn6").click(function(){
		pushback($("#extract_partner"),"../extract_partner.php")
	})
	//小额礼包
	$("#t_btn").click(function(){
		pushback($("#Money_partner"),"../money.php")
	})

	//模拟
	// 模拟充值模块ajax
	$("#mn_but").click(	function(){
			pushback($("#simulator"),"../do_simulator.php")
		})
	// 执行命令模块ajax
	$("#zx_but").click(	function(){
			pushback($("#execute"),"../do_execute.php")
		})

//	GM功能
	//服务器
	function server_1(aa){
		$.ajax({
			type:"post",
			url:"../query_area_servers.php",
			dataType:"json",
			success:function(e){
				for(var key in e){
					var tim=e[key]
					for(var i=0; i<tim.length;i++){
						 var id=tim[i].id
						 lable='<tr ><td>'+id+'</td><td>'+tim[i].name+'</td> <td>'+tim[i].areaname+'</td><td>'+tim[i].capa+'</td></tr>'
						aa.append(lable)
					}
				}
			}
		})
	}
	server_1($(".s_table"))
//	CDK生成
	$("#w_btn").click(function(){
		pushback($("#cdkey_generate"),"../cdkey_generate.php")
	})
//	CDK查询
	$("#q_btn").click(function(){
		pushback($("#cdkey_select"),"../cdkey_select.php")
	})
//	聊天查询(带分页）
	//在线查询ajax
	$("#btn7").click(function(){
		$.ajax({
			type:"post",
			url:"../query_now_chat.php",
			dataType:"json",
			data:$("#query_now_chat").serialize(),
			success:function(e){

				 var data = [];
				    for(var i=0;i<e.length;i++)
				    {
				        data[i] =[e[i].SortId,e[i].Timestamp,e[i].PlayerGuid,e[i].PlayerName,e[i].ChannelId,e[i].Content];
					}
				    var cs = new table({
				        "tableId":"cs_table",    //必须 表格id
				        "headers":["主键id","时间","PlayerId","玩家名字","频道","内容"],   //必须 thead表头
				        "data":data,         //必须 tbody 数据展示
				        "displayNum": 20,    //必须   默认 10  每页显示行数
				        "groupDataNum":1,     //可选    默认 10  组数
				        "display_tfoot":true, // true/false  是否显示tfoot --- 默认false
				        "bindContentTr":function(){ //可选 给tbody 每行绑定事件回调
				            this.tableObj.find("tbody").on("click",'tr',function(e){
				                return false;
				                var tr_index = $(this).data("tr_index");        // tr行号  从0开始
				                var data_index = $(this).data("data_index");   //数据行号  从0开始
				            })
				        },
				        sort:true,    // 点击表头是否排序 true/false  --- 默认false
				        search:true   // 默认为false 没有搜索
				    });
				},
			error:function(e){
				alert("执行错误")
			}
		})
	})
	//历史查询ajax
	$("#btn8").click(function(){
		$.ajax({
			type:"post",
			url:"../query_history_chat.php",
			dataType:"json",
			data:$("#history_chat").serialize(),
			success:function(e){

				 var data = [];
				    for(var i=0;i<e.length;i++)
				    {
				        data[i] =[e[i].SortId,e[i].Timestamp,e[i].PlayerGuid,e[i].PlayerName,e[i].ChannelId,e[i].Content];
					}
				    var cs = new table({
				        "tableId":"ds_table",    //必须 表格id
				        "headers":["主键id","时间","id","玩家名字","频道","内容"],   //必须 thead表头
				        "data":data,         //必须 tbody 数据展示
				        "displayNum": 20,    //必须   默认 10  每页显示行数
				        "groupDataNum":1,     //可选    默认 10  组数
				        "display_tfoot":true, // true/false  是否显示tfoot --- 默认false
				        "bindContentTr":function(){ //可选 给tbody 每行绑定事件回调
				            this.tableObj.find("tbody").on("click",'tr',function(e){
				                return false;
				                var tr_index = $(this).data("tr_index");        // tr行号  从0开始
				                var data_index = $(this).data("data_index");   //数据行号  从0开始
				            })
				        },
				        sort:true,    // 点击表头是否排序 true/false  --- 默认false
				        search:true   // 默认为false 没有搜索
				    });
				},
				error:function(e){
					alert("执行错误")
				}
		})
	})


	//角色查询ajax
	$("#btn1").click(function(){
		$.ajax({
			type:"post",
			url:"../query_player.php",
			dataType:"html",
			data:$("#query_player").serialize(),
			success:function(e){
				$(".w_text").html(e)
			},
			error:function(e){
				alert("错误")
			}
		})
	})

	$(".w_button").click(function(){
		$(".t_table").tableExport({formats:["xlsx"]});
	})
	$(".b_button").click(function(){
		$(".b_table").tableExport({formats:["xlsx"]});
	})
	//管理员
	$("#Admin_btn").click(function(){
		window.location.href="../home/regist.html"

	})
	//运营活动
	//添加
	//顶级招募
	function add(aa,bb){
		$(aa).click(function(){
		for (var i=0;i<10;i++){
			var tab='<tr align="center"><td width="18%"><input  type="text" name="ckb'+i+'"/></td><td width="18%" class="w_wid"><input  type="text"  name="ckab'+((i*3)+1)+'"/><input  type="text" name="ckab'+((i*3)+2)+'"/><input  type="text" name="ckab'+((i*3)+3)+'"/></td><td width="18%" class="w_wid"><input  type="text" name="ckeb'+((i*3)+1)+'"/><input  type="text" name="ckeb'+((i*3)+2)+'"/><input  type="text" name="ckeb'+((i*3)+3)+'"/></td><td width="18%"><input class="w_check" type="checkbox" name="ckbc"/></td></tr>'
			bb.append(tab)
			}
		})
	}
	//登录
	function addlogin(aa,bb){

		$(aa).click(function(){
		for (var i=0;i<7;i++){
			var tab='<tr align="center"><td width="18%"><input  type="text" name="ckb'+i+'"/></td><td width="18%" class="w_wid"><input  type="text"  name="ckab'+((i*3)+1)+'"/><input  type="text" name="ckab'+((i*3)+2)+'"/><input  type="text" name="ckab'+((i*3)+3)+'"/></td><td width="18%" class="w_wid"><input  type="text" name="ckeb'+((i*3)+1)+'"/><input  type="text" name="ckeb'+((i*3)+2)+'"/><input  type="text" name="ckeb'+((i*3)+3)+'"/></td><td width="18%"><input class="w_check" type="checkbox" name="ckbc"/></td></tr>'
			bb.append(tab)
			}
		})

	}
	//打折商店
	function add1(aa,bb){

		$(aa).click(function(){
		var tab='<tr align="center"><td width="16%"><input  type="text" name="ckb"/></td><td width="16%" class="w_wid wid"><input  type="text" name="ckeb1"/></td><td width="16%" class="w_wid wid"><input  type="text" name="ckeb2"/></td><td width="16%" class="w_wid wid"><input  type="text" name="ckeb3"/></td><td width="16%" class="w_wid wid"><input  type="text" name="ckeb4"/></td><td width="16%"><input class="w_check" type="checkbox" name="ckbc"/></td></tr>'

			bb.append(tab)
		})
	}
	//单笔充值
	function add2(aa,bb){

		$(aa).click(function(){
			for(var i=0;i<8;i++){
				var tab='<tr align="center"><td width="16%"><input  type="text"  name="ckb'+i+'"/></td><td width="16%" class="w_wid "><input  type="text" name="ckab'+((i*3)+1)+'"/><input  type="text" name="ckab'+((i*3)+2)+'"/><input  type="text" name="ckab'+((i*3)+3)+'"/></td><td width="16%" class="w_wid "><input  type="text" name="ckeb'+((i*3)+1)+'"/><input  type="text" name="ckeb'+((i*3)+2)+'"/><input  type="text" name="ckeb'+((i*3)+3)+'"/></td><td width="16%" class="w_wid q_wid"><input  type="text" name="ckcb"/></td><td width="16%" ><input class="w_check" type="checkbox" name="ckbc"/></td></tr>'
				bb.append(tab)
			}
		})
	}
	//累积充值
	function addrecharge(aa,bb){
		$(aa).click(function(){
		for (var i=0;i<8;i++){
			var tab='<tr align="center"><td width="18%"><input  type="text" name="ckb'+i+'"/></td><td width="18%" class="w_wid"><input  type="text"  name="ckab'+((i*3)+1)+'"/><input  type="text" name="ckab'+((i*3)+2)+'"/><input  type="text" name="ckab'+((i*3)+3)+'"/></td><td width="18%" class="w_wid"><input  type="text" name="ckeb'+((i*3)+1)+'"/><input  type="text" name="ckeb'+((i*3)+2)+'"/><input  type="text" name="ckeb'+((i*3)+3)+'"/></td><td width="18%"><input class="w_check" type="checkbox" name="ckbc"/></td></tr>'
			bb.append(tab)
			}
		})
	}
	//登录
	addlogin($('#activity-1 .w_add'),$('#activity-1 .btable tbody'))
	//累积充值
	addrecharge($('#activity-2 .w_add'),$('#activity-2 .btable tbody'))
	//打折商店
	add1($('#activity-3 .w_add'),$('#activity-3 .btable tbody'))
	//单笔充值
	add2($('#activity-4 .w_add'),$('#activity-4 .btable tbody'))
	//顶级招募
	add($('#activity-6 .w_add'),$('#activity-6 .btable tbody'))

	//删除
	function remove(aa){
		$(aa).click(function(){
			for(var i=0;i<$(".w_check").length;i++){
				if($(".w_check").eq(i).prop("checked")==true){
					var len=$("input[type=checkbox]:checked()")
					$(len).parent().parent().remove()
				}
			}
		})
	}
	//登录
	remove($('#activity-1 .w_remove'))
	//累积充值
	remove($('#activity-2 .w_remove'))
	//打折商店
	remove($('#activity-3 .w_remove'))
	//单笔充值
	remove($('#activity-4 .w_remove'))
	//顶级招募
	remove($('#activity-6 .w_remove'))
	
		
	
	
	//时间插件
	$('.form_datetime').datetimepicker({
        //language:  'fr',
        weekStart: 1,
		todayBtn:  1,
		autoclose: 1,
		todayHighlight: 1,
		startView: 2,
		forceParse: 0,
	    showMeridian: 1,
	  	format: "yyyy-MM-dd hh:ii"
    });
	$('.form_date').datetimepicker({
        language:  'fr',
        weekStart: 1,
        todayBtn:  1,
		autoclose: 1,
		todayHighlight: 1,
		startView: 2,
		minView: 2,
		forceParse: 0
    });
	$('.form_time').datetimepicker({
        language:  'fr',
        weekStart: 1,
        todayBtn:  1,
		autoclose: 1,
		todayHighlight: 1,
		startView: 1,
		minView: 0,
		maxView: 1,
		forceParse: 0
    });


})


