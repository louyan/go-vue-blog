<template>
    <section class="post-list right-container">
        <back-menu></back-menu>
        <back-header></back-header>
        <post-table :post-list="postList" @delete-post='changePostTotal'></post-table>
        <pagination :page-size="pageSize"
                    :total="total"
                    :current-page="page"
                    @change-page="changePage">
        </pagination>
    </section>
</template>

<script>
    import api from '../fetch/api';
    import Menu from '../components/Menu';
    import Header from '../components/Header';
    import PostTable from '../components/PostTable';
    import Pagination from '../components/Pagination';
    export default {
        components: {
            'back-menu': Menu,
            'back-header': Header,
            PostTable,
            Pagination
        },
        data() {
            return {
                postList: [],
                page:1,
                pageSize:10,
                total:0
            }
        },
        created(){
            this.getPostTotal();
            this.getPostList();
        },
        methods: {
            getPostList: async function () {
                let res = await api.getPostList(this.page, this.pageSize);
                console.log(res)
                console.log(res.success)
                if (res.success === 1) {
                    this.postList = res.posts;
                    console.log(res.posts)
                }
            },
            getPostTotal: async function () {
                let res = await api.getPostTotal();
                if (res.success === 1) {
                    this.total = res.total;
                    console.log( this.total)
                }
            },
            changePage(newPage) {
                this.page = newPage;
                this.getPostList();
            },
            changePostTotal() {
                this.getPostTotal();
                this.getPostList();
            }
        },

    }
</script>

<style scoped>

</style>