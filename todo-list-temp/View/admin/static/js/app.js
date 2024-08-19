$(document).ready(function(){
	$(function(){
		$(".complete").on("click",function(){
			var id=$(this).attr("rel")
			$.ajax({
				url:"/admin/todolist/done/"+id,
				type:"GET",
				success:function(){
					window.location.reload();
				}
			});
		});
	});
	$(function(){
		$(".delete").on("click",function(){
			var id=$(this).attr("rel")
			$.ajax({
				url:"/admin/todolist/delete/"+id,
				type:"GET",
				success:function(){
					window.location.reload();
				}
			});
		});
	});
});