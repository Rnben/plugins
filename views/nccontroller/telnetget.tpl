<!DOCTYPE html>
<html>
  
  <head>
    <title>T</title>
    <meta name="viewport" content="width=device-width, initial-scale=1" charset="utf-8">
    <script src="https://cdnjs.cloudflare.com/ajax/libs/vue/2.2.6/vue.js"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/vue-resource/1.3.1/vue-resource.min.js"></script>
    <!-- jQuery (Bootstrap 插件需要引入) -->
    <script src="http://cdn.static.runoob.com/libs/jquery/2.1.1/jquery.min.js"></script>
    <link rel="stylesheet" href="http://cdn.static.runoob.com/libs/bootstrap/3.3.7/css/bootstrap.min.css">  
    <!-- 包含了所有编译插件 -->
    <script src="http://cdn.static.runoob.com/libs/bootstrap/3.3.7/js/bootstrap.min.js"></script>

    <script type="text/javascript">window.onload = function() {
        Vue.http.options.emulateHTTP = true;
        var vm = new Vue({
          el: '#box',
          data: {
            //requrl: 'http://10.78.177.47:8082',
            requrl: 'http://20.26.28.83:8082',
            telnetinfo:{}
          },
          methods: {
            getinfo: function() {
              //发送get请求
              this.$http.get(this.requrl+'/version').then(function(res) {
                alert(res.body);
              },
              function() {
                alert('请求失败'); //失败处理
              });
            },

            telnet: function() {
              //发送post请求
              var vm=this
              vm.$http.post(vm.requrl+'/telnet',vm.telnetinfo).then(function(res) {
                alert(res.body);
              },
              function() {
                alert('请求失败'); //失败处理
              });
            },
          }
        });
      }</script>
	  <style type="text/css">
		  .color{color:#FF2200;}
	  </style>
  </head>
  
  <body>
    <div id="box" class="container">
      <div class="">
        <h1 class="color text-center">主机连通性测试</h1>
      </div>
      <div class="input-group">
        <div class="">
		  <span class="input-group-btn"><button class="btn btn-default" type="button">源地址</button></span>
		  <input class="form-control" type="text" name="srcip" v-model="telnetinfo.srcip" placeholder="192.168.1.100" aria-describedby="basic-addon1"/>
        </div>          
        <div class="">
          <span class="input-group-btn"><button class="btn btn-default" type="button">目的地址</button></span>
          <input class="form-control" type="text" name="desip" v-model="telnetinfo.desip" placeholder="192.168.1.101"/>
        </div>
        <div class="">
          <span class="input-group-btn"><button class="btn btn-default" type="button">目的端口</button></span>
          <input class="form-control" type="text" name="desport" v-model="telnetinfo.desport" placeholder="8080"/>
        </div>
        <div>
          <input class="btn btn-lg btn-info btn-block" type="submit" @click="telnet()" value="Telnet"/>
        </div>
      </div>
    </div>
  </body>

</html>
