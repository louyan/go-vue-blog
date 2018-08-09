import Vue from 'vue'
import App from './App.vue'
import router from './router'
import axios from 'axios'
// import VueSocketio from 'vue-socket.io';
// import socketio from 'socket.io-client';

import MessageBox from './components/MessageBox/index';
import Message from './components/Message/index';
Vue.use(MessageBox);
Vue.use(Message);

// Vue.use(VueSocketio, socketio('http://localhost:8081/v1', {
//     path: '/testsocketiopath'
// }));

Vue.config.productionTip = false

axios.defaults.baseURL = 'http://127.0.0.1:8081/v1'
// axios.defaults.headers.post['Content-Type'] = "application/json"
axios.defaults.headers.post['Content-Type'] = 'application/x-www-form-urlencoded'
// Vue.prototype.$axios = axios
// axios.post('/user/login',{
//     params:{
//         Name:"admin",
//         Password:"123456"
//     }
// }).then((res)=>{
//     console.log(res.data);
// }).catch((err)=>{
//     console.log(err);
// })

axios.interceptors.response.use(
    response => {
        return response;
    },
    error => {
        if (error.response.status === 401){
            Vue.prototype.$msgBox.showMsgBox({
                title: '错误提示',
                content: '您的登录信息已失效，请重新登录',
                isShowCancelBtn: false
            }).then((val) =>{
                router.push('/login');
            }).catch(()=>{
                console.log('cancel');
            });
        }else {
            Vue.prototype.$message.showMessage({
                type: 'error',
                content: '系统出现错误'
            });
        }

        return Promise.resolve(error);
    }
);

new Vue({
  render: h => h(App),
    router,
}).$mount('#app')
