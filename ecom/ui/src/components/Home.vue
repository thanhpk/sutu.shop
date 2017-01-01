<template src="./Home.jade" lang="jade">
</template>

<script>
	import { carousel } from 'vue2-strap';
	import PageFooter from './Footer.vue';
	import $ from 'jquery';
	import FunButton from './FunButton.vue';
	import Cart from './Cart.vue';
	export default {
	components: { carousel, PageFooter, FunButton, cart: Cart },
	mounted: function() {

	}
	}
</script>

<style>
	@import url('~assets/css/feather.css');

	.small-product {
	position: relative;
	margin-top: 20px;
	}
	
	.small-product img {
	object-fit: cover;
	position: absolute;
	width: 40%;
	height: 100px;
	border-radius: 5px;
	
	}

	.logo-text {
		font-weight: 300;
		}
	
	.small-product .content {
	display: block;
	margin-left: 45%;
	}

	.small-product h4 {
	font-size: 16px;
	font-weight: 500;
	text-transform: capitalize;
	}

	.product {
	text-align: center;
	margin: auto;
	max-width: 400px;
	
	}

	.product .icon-bag {
	position: relative;
	top: -3px;
	margin-right: 10px;
	}

	.gal-item {
	width:	16.66%;
	display: inline-block;
	}

	.gal-item img {
		width: 100%;
		height: auto;
		}
	.product code:hover {
	color: #444;
	transition: background-color 0.1s ease-out, color 0.2s ease-out;
	background: #ffd32c;
	
	}
	

	
	.product code {
  transition: background-color 0.1s ease-out, color 0.2s ease-out;
	font-size: 15px;
	border-radius: 3px;
	display: inline-block;
	height: 38px;
	font-family: 'Myriad Set Pro';
	font-weight: 400;
	color: #f2f2f2;
	background-color: #444;

  padding-left: 15px;
  padding-right: 15px;
  padding-top: 10px;

	}

	.product .dot:before {
  content: "\002a";
	}
	.product .dot {

	position: relative;
  top: 1px;
  display: inline-block;
  font-family: 'Glyphicons Halflings';
  font-style: normal;
  font-weight: 400;
  line-height: 1;
  -webkit-font-smoothing: antialiased;
	
	font-size: 9px;
  vertical-align: super;
  /* margin-top: 10px; */
  /* line-height: 1; */
  top: 5px;
  left: 0px;

	}
	
	.product h4 {
	margin-top: 20px;
	font-weight: 400;
	text-transform: capitalize;
	font-size: 19px;
	margin-bottom: 15px;
	}
	.product-img {
	border-radius: 15px;
	height: 250px;
	background-origin: content-box;
	background-position: center;
	background-size: cover;
	}
	
	.menu-skincare {
	background: url('/static/images/skincare.jpg');
	}
	
	.menu-makeup {
	background: url('/static/images/makeup.jpg');
	}

	.menu-hair {
	background: url('/static/images/hair.jpg');
	}
	.menu {
	margin-top: 15px;
	height: 100px;

	-webkit-background-size:cover;
	-moz-background-size:cover;
	-o-background-size:cover;
	background-size:cover;
	backgorund-repeat: no-repeat;
	}

	.menu a:hover {
	text-decoration: none;
  text-shadow: 1px 1px 1px rgba(0,0,0,0.1);
  color: rgba(255,255,255,0.9);
	}

	.menu a {
	text-decorator: none;
	width: 100%;
	display: block;
  font-weight: 300;
  padding: 2%;
  font-size: 66px;
  letter-spacing: 7px;
  text-shadow: 0px 0px 10px rgba(255,255,255,0.9);
  color: rgba(255,255,255,0.1);
	
  -webkit-background-size: 100%;
  -moz-background-size: 100%;
  -o-background-size: 100%;
  background-size: 100%;
  background-clip: padding-box;
  -webkit-transition: all 0.3s ease-in-out;
  -moz-transition: all 0.3s ease-in-out;
  -o-transition: all 0.3s ease-in-out;
  -ms-transition: all 0.3s ease-in-out;
  transition: all 0.3s ease-in-out;
	}

</style>



