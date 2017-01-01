import Vue from 'vue';

import BootstrapVue from 'bootstrap-vue';

import VueRouter from  'vue-router';
import Home from './components/Home';
import BuyLayout from './components/BuyLayout.vue';
import Category from './components/Category.vue';
import Product from './components/Product.vue';
import FbScript from './components/FbScript.vue';

Vue.use(BootstrapVue);
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
														 
/* eslint-disable no-new */
const app = new Vue({
	router: router,
	components: {
		fbscript: FbScript
	}
}).$mount('#app');
