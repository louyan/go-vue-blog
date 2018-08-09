import async from './fetch';
const baseUrl = '/admin';

export default {
    getPostById(id){
        return async(`${baseUrl}/${id}`);//getPostById/
    },
    addPost(params){
        return async(`${baseUrl}/addPost`, params, 'post');
    },
    updatePost (id, params) {
        return async(`${baseUrl}/updatePost/${id}`, params, 'post');
    },
    getPostList (page = 1, pageNum = 10) {
        return async(`${baseUrl}/getPostList`, {
            page: page,
            pageNum: pageNum
        });
    },
    // getPostList (page = 1, pageNum = 10) {
    //     return async(`${baseUrl}/getPostList`, {
    //         page: page,
    //         pageNum: pageNum
    //     });
    // },
    getPostTotal () {
        return async(`${baseUrl}/getPostTotal`);
    },
    offlinePost (id) {
        return async(`${baseUrl}/offlinePost/${id}`, {}, 'put');
    },
    publishPost (id) {
        return async(`${baseUrl}/publishPost/${id}`, {}, 'put');
    },
    deletePost (id) {
        return async(`${baseUrl}/deletePost/${id}`, {}, 'delete');
    },

};