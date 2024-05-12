var vue = new Vue({
    el: "#app",
    // data: {
    //     file: {id:"",fileName:"aaa",filePath:"",fileSize:"",lastUpdateTime:"",status:""},
    //     fileList: []
    // },
    data: {
        file: {id:"",community_id:"",community_name:"aaa",introduction:"",create_time:"",update_time:""},
        fileList: [],
        fileColumns: [],
        div_heads_message:""
    },
    methods: {
        findAll: function () {
            var _this = this;
            axios.get("http://127.0.0.1:8080/api/v1/community").then(function (response) {
                _this.fileList = response.data;
                console.log(_this.fileList);
            }).catch(function (err) {
                console.log(err);
            });
        },
        findById: function (fileid) {
            var _this = this;
            axios.get("/file/findById", {
                params: {
                    id: fileid
                }
            }).then(function (response) {
              _this.file = response.data;
                $('#myModal').modal("show");
            }).catch(function (err) {
            });

        },
        update: function (file) {
            var _this = this;
            axios.post("/file/update",_this.file).then(function (response) {
                _this.findAll();
            }).catch(function (err) {
            });
        },

        getAnalyzeData:function (){
             $('#content-section').load('analyze_content_section.html');
//             $('#content-wrapper-body').load('login.html');
        },

        getCompareData:function (){
            console.log("-------------section loading-------");
            $('#content-section').load('compare_content_section.html');
             console.log("-------------section loeaded-------");
        }

    },

    created(){
        this.findAll();
    }
});