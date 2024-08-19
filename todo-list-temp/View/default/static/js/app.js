$(document).ready(function(){
    $(function() {
		$.ajax({
			url:"/api/v1/todoView/",
			type:"get",
			dataType:"json",
			success:function(data){
				console.log(data);
			}
		});
	});
});