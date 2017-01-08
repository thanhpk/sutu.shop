import Vue from 'vue';

import VueRouter from  'vue-router';
import Home from './components/Home.vue';
import BuyLayout from './components/BuyLayout.vue';
import Category from './components/Category.vue';
import Product from './components/Product.vue';
import FbScript from './components/FbScript.vue';

import '../static/bower_components/animate.css/animate.min.css';
import '../static/fonts/flaticon.css';
import '../static/app.css';

Vue.use(VueRouter);

const router = new VueRouter({
	routes: [
		{
			path: '/',
			component: Home
		}, {
			path: '/makeup', component: BuyLayout,
			children: [{ path: '', component: Category }]
		}, {
			path: '/product', component: BuyLayout,
			children: [{ path: '', component: Product }]
		}, {
			path: '/product/:id', component: BuyLayout,
			children: [{ path: '', component: Product }]
		},
	]
});
														 
new Vue({
	router: router,
	components: {
		fbscript: FbScript
	}
}).$mount('#app');
