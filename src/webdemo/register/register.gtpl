<html>
<head>
<script type="text/javascript">  
    function displayAlert(){  
        setInterval(go, 1000);  
    };  
    var x=3; //利用了全局变量来执行  
    function go(){  
        x--;  
        if(x>0){  
        document.getElementById("sp").innerHTML= x + 秒后自动跳转到登录页面;  //每次设置的x的值都不一样了。  
        }else{  
        location.href='res.html';  
        }  
    }  
</script>
</head>
<body>
<form action="/register" method="post">
用户名:<br/>
<input type="text" name="username"/>
<br/><br/>
密码:<br/>
<input type="password" name="pwd"/>
<br/><br/>
确认密码:<br/>
<input type="password" name="pwd2"/>
<br/>
<input type="submit" value="注册"/>
<input type="button" onclick="displayAlert()" value="测试"/>
<sp id="sp"/>
</form>
</body>
</html>