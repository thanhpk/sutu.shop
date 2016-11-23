import Vue from 'vue';
import VueRouter from  'vue-router';
import Home from './components/Home';
import Hello from './components/Hello';

Vue.use(VueRouter);

const router = new VueRouter({
	routes: [
		{
			path: '/',
			component: Home
		},
		{
			path: '/hello',
			component: Hello
		}
	]
});
														 
/* eslint-disable no-new */
const app = new Vue({
	router: router
}).$mount('#app');
