<template>
	<footer class="container mtbig">
		<div class="row">

<div class="col-sm-5 mtmed">
<img class="logo" src="static/images/logo-gray.png"/>
				<h4>Về chúng tôi</h4>
<p>Shop bán quần áo sutu.shop</p>
<p><b>Dia chi:</b> Số 68B ngõ 204 Trần Duy Hưng Phường Trung Hòa, Quận Cầu Giấy, Hà Nội</p>
<p><b>Dien thoai:</b> 0167 861 6688</p
<p><b>Email:</b> sutu.shop@gmail.com</p>
<p><a href="//facebook.com/sutu.shop"><img class="social" src="static/images/facebook.png" /> <a href="//instagram.com/sutu.shop"><img class="social" src="static/images/instagram.png"/></a></p>
			</div>

			<div class="col-sm-3 mtmed">
<h4>Tai khoan</h4>
<p><a>Dang ky</a></p>
<p><a>Dang nhap</a></p>
<h4>Quyen loi khach hang</h4>
<router-link class="footer-link" to="/return">Chính sách đổi trả</router-link>
			<p><router-link class="footer-link" to="/privacy">Quyền riêng tư</router-link></p>
<p><router-link class="footer-link" to="/term">Dieu khoan</router-link></p>
			</div>

			<div class="col-sm-4 mtmed">
<img class="seal" src="https://upload.wikimedia.org/wikipedia/commons/6/6d/Norton_av_logo.png"/>
<img class="seal" src="http://online.gov.vn/Images/dathongbao.png" />
			</div>

		</div>
		
<div class="row"><div class="col-sm-12 mtmed mbmed">
 <p class="pull-right">© 2016 sutu.shop</p>

</div></div>
	</footer>
</template>
<script>
	
</script>
<style>
	.seal {
	height: 50px;
	display: block;
	}
	.footer-link {
	text-transform: uppercase;
	margin-right: 20px;
	}
	.logo {
			position: absolute;
	width: 100px;
	opacity: 0.3;
	}
	h4 {
	text-transform: uppercase;
	font-weight: 500;
	font-size: 14px;
	margin-bottom: 2px;
	}
	p {
	margin-bottom: 2px;
	}
	.label {
	font-size: 12px;
	vertical-aligh: text-top;
	font-weight: 100;
	}
	footer {
	border-top: 1px solid lightgray;
	}
.social {
		width: 20px;
		height: 20px;
}	
	</style>
