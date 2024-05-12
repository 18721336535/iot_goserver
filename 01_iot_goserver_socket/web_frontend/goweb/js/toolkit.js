var vue = new Vue({
    el: "#toolkit-app",
    data: {
        file: {id:"",community_id:"",community_name:"aaa",introduction:"",create_time:"",update_time:""},
        fileList: [],
        fileColumns: [],
        div_heads_message:""
    },

    methods: {
        findAll: function () {
            var _this = this;
            axios.get("/api/v1/community").then(function (response) {
                _this.fileList = response.data.data;
                console.log(_this.fileList);
                console.log(_this.fileList.data);
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
        getFileHeads: function () {
            var _this = this;
            var filePath = $("#search-path-id").val();
            var field_key_value = $("#filed-search-key-value-id").val();
            let paramsDto = {filePath:filePath,fieldKeyValue:field_key_value}
            axios.post("/file/getFileHeads",paramsDto).then(function (response) {
                  var arr = [];
                  arr = response.data;
                  var headsHtml = "<table><tr>";
                  for(var i = 0; i < arr.length; i++){
                        headsHtml = headsHtml+"<td>"+arr[i]+"</td>";
                  }
                  headsHtml = headsHtml+"</tr></table>";
                  _this.div_heads_message = headsHtml;
                  console.log(_this.fileList);
            }).catch(function (err) {
                  console.log(err);
            });
        },
        findDataByFiledNameAndValue: function () {
            var _this = this;
            var filePath = $("#search-path-id").val();
            var field_key_value = $("#filed-search-key-value-id").val();
            let paramsDto = {filePath:filePath,fieldKeyValue:field_key_value}
            axios.post("/file/findDataByFiledNameAndValue",paramsDto).then(function (response) {
                  _this.fileList = response.data;
                  console.log(_this.fileList);
            }).catch(function (err) {
                  console.log(err);
            });
        },
        update: function (file) {
            var _this = this;
            axios.post("/file/update",_this.file).then(function (response) {
                _this.findAll();
            }).catch(function (err) {
            });
        }
    },

    created(){
        this.findAll();
    }
});