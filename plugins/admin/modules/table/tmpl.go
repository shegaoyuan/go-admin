package table

var tmpls = map[string]string{"choose_table_ajax":`{{define "choose_table_ajax"}}
        NProgress.start();
        let info_table = $("tbody.fields-table");
        info_table.find("tr").remove();
        let tpl = $("template.fields-tpl").html();
        for (let i = 0; i < data.data[0].length; i++) {
            info_table.append(tpl);
        }
        let trs = info_table.find("tr");
        for (let i = 0; i < data.data[0].length; i++) {
            $(trs[i]).find('.field_head').val(data.data[0][i]);
            $(trs[i]).find('.field_name').val(data.data[1][i]);
            $(trs[i]).find('select.field_db_type').val(data.data[2][i]).select2();
        }
        let form_table = $("tbody.fields_form-table");
        form_table.find("tr").remove();
        let tpl_form = $("template.fields_form-tpl").html();
        for (let i = 0; i < data.data[0].length; i++) {
            form_table.append(tpl_form);
        }
        let trs_form = form_table.find("tr");
        let pk = $(".pk").val();
        for (let i = 0; i < data.data[0].length; i++) {
            $(trs_form[i]).find('.field_head_form').val(data.data[0][i]);
            $(trs_form[i]).find('.field_name_form').val(data.data[1][i]);
            $(trs_form[i]).find('input.field_canedit').iCheck("check");
            if (!(data.data[1][i] === pk || (pk === "" && data.data[1][i] === "id"))) {
                $(trs_form[i]).find('input.field_canadd').iCheck("check");
            }
            if (data.data[1][i] === "created_at" || data.data[1][i] === "updated_at") {
                $(trs_form[i]).find('select.field_display').val("1").select2();
            }
            $(trs_form[i]).find('select.field_db_type_form').val(data.data[2][i]).select2();
            $(trs_form[i]).find('select.field_form_type_form').val(data.data[3][i]).select2();
        }
        $(".hide_filter_area.ga_checkbox").bootstrapSwitch('state', true);
        $(".hide_new_button.ga_checkbox").bootstrapSwitch('state', true);
        $(".hide_export_button.ga_checkbox").bootstrapSwitch('state', true);
        $(".hide_edit_button.ga_checkbox").bootstrapSwitch('state', true);
        $(".hide_pagination.ga_checkbox").bootstrapSwitch('state', true);
        $(".hide_delete_button.ga_checkbox").bootstrapSwitch('state', true);
        $(".hide_detail_button.ga_checkbox").bootstrapSwitch('state', true);
        $(".hide_filter_button.ga_checkbox").bootstrapSwitch('state', true);
        $(".hide_row_selector.ga_checkbox").bootstrapSwitch('state', true);
        $(".hide_query_info.ga_checkbox").bootstrapSwitch('state', true);
        $(".filter_form_layout.ga_checkbox").bootstrapSwitch('state', true);
        $(".hide_continue_edit_check_box.ga_checkbox").bootstrapSwitch('state', true);
        $(".hide_reset_button.ga_checkbox").bootstrapSwitch('state', true);
        $(".hide_continue_new_check_box.ga_checkbox").bootstrapSwitch('state', true);
        $(".hide_back_button.ga_checkbox").bootstrapSwitch('state', true);

        let detail_table = $("tbody.fields_detail-table");
        detail_table.find("tr").remove();

        NProgress.done();
{{end}}`,"generator":`{{define "generator"}}
    <script>
        $(function () {
            let pack = localStorage.getItem("{{index . "prefix"}}package");
            let pk = localStorage.getItem("{{index . "prefix"}}pk");
            let path = localStorage.getItem("{{index . "prefix"}}path");
            if (pack !== "") {
                $(".package").val(pack);
            }
            if (pk !== "") {
                $(".pk").val(pk);
            }
            if (path !== "") {
                $(".path").val(path);
            }

            let save_table_list_str = localStorage.getItem("{{index . "prefix"}}save_table_list");
            if (save_table_list_str && save_table_list_str !== "") {
                addTableToList(JSON.parse(save_table_list_str));
            }
        });

        function getLis() {
            return $("li.list-group-item.list-group-item-action")
        }

        function addTableToList(save_table_list) {
            let list_group = $(".list-group.save_table_list");
            getLis().remove();
            for (let i = save_table_list.length - 1; i > save_table_list.length - 6 && i > -1; i--) {
                let new_li = "<li class='list-group-item list-group-item-action'>" + save_table_list[i] + "</li>";
                list_group.append(new_li);
            }
            list_group.show();
            getLis().on("click", restoreTableData);
        }

        $(".nav.nav-tabs li a").on("click", function () {
            let href = $(this).attr("href");
            let trs = $(".list-group-item.list-group-item-action");
            if (trs.length > 0) {
                if (href === "#tab-form-0") {
                    $(".list-group.save_table_list").show();
                }
                if (href === "#tab-form-1" || href === "#tab-form-2" || href === "#tab-form-3") {
                    $(".list-group.save_table_list").hide();
                }
            }
        });

        $(".btn-group.pull-right .btn.btn-primary").on("click", function () {
            let pack = $(".package").val();
            let pk = $(".pk").val();
            let path = $(".path").val();
            if (pack !== "") {
                localStorage.setItem("{{index . "prefix"}}package", pack);
            }
            if (pk !== "") {
                localStorage.setItem("{{index . "prefix"}}pk", pk);
            }
            if (path !== "") {
                localStorage.setItem("{{index . "prefix"}}path", path);
            }
            let table = $("select.table").val();
            if (table && table !== "") {
                let save_table_list = [];
                let save_table_list_str = localStorage.getItem("{{index . "prefix"}}save_table_list");
                if (save_table_list_str && save_table_list_str !== "") {
                    save_table_list = JSON.parse(save_table_list_str);
                }
                let table_index = save_table_list.indexOf(table);
                if (table_index !== -1) {
                    save_table_list.splice(table_index, 1);
                }
                save_table_list.push(table);
                localStorage.setItem("{{index . "prefix"}}save_table_list", JSON.stringify(save_table_list));
                localStorage.setItem("{{index . "prefix"}}save_table_" + table, getItemObjData());

                addTableToList(save_table_list);
            }
        });

        function restoreTableData() {
            NProgress.start();
            let data_str = localStorage.getItem("{{index . "prefix"}}save_table_" + $(this).html());
            if (data_str && data_str !== "") {
                let data = JSON.parse(data_str);
                $(".package").val(data.package);
                $(".pk").val(data.pk);
                $(".path").val(data.path);
                $("select.conn").val(data.conn).select2();
                conn_req_refresh($("select.table"), false, "select");
                setTimeout(function () {
                    $("select.table").val(data.table).select2();
                }, 2000);

                $(".extra_import_package").val(data.extra_import_package).select2();
                $(".table_title").val(data.table_title);
                $(".table_description").val(data.table_description);
                $(".form_title").val(data.form_title);
                $(".form_description").val(data.form_description);
                $(".detail_title").val(data.detail_title);
                $(".detail_description").val(data.detail_description);
                $("select.detail_display").val(data.detail_display).select2();
                if (data.detail_display !== "0") {
                    $("label[for='detail_title']").parent().show();
                    $("label[for='detail_description']").parent().show();
                    $("label[for='fields_detail']").parent().show();
                }

                if (extra_codeeditor && data.extra_code && data.extra_code !== "") {
                    extra_codeeditor.setValue(decodeURIComponent(data.extra_code));
                }

                let info_table = $("tbody.fields-table");
                info_table.find("tr").remove();
                let tpl = $("template.fields-tpl").html();
                for (let i = 0; i < data.infos.length; i++) {
                    info_table.append(tpl);
                }

                let trs = info_table.find("tr");
                for (let i = 0; i < trs.length; i++) {
                    $(trs[i]).find('.field_head').val(data.infos[i][0]);
                    $(trs[i]).find('.field_name').val(data.infos[i][1]);
                    checkItemSwitch($(trs[i]).find('input.field_filterable'), data.infos[i][2]);
                    checkItemSwitch($(trs[i]).find('input.field_sortable'), data.infos[i][3]);
                    checkItemSwitch($(trs[i]).find('input.field_hide'), data.infos[i][4]);
                    checkItemSwitch($(trs[i]).find('input.info_field_editable'), data.infos[i][5]);
                    $(trs[i]).find('select.field_db_type').val(data.infos[i][6]).select2();
                }

                let form_table = $("tbody.fields_form-table");
                form_table.find("tr").remove();
                let tpl_form = $("template.fields_form-tpl").html();
                for (let i = 0; i < data.forms.length; i++) {
                    form_table.append(tpl_form);
                }

                let trs_form = form_table.find("tr");
                for (let i = 0; i < trs_form.length; i++) {
                    $(trs_form[i]).find('.field_head_form').val(data.forms[i][0]);
                    $(trs_form[i]).find('.field_name_form').val(data.forms[i][1]);
                    checkItemSwitch($(trs_form[i]).find('input.field_canedit'), data.forms[i][2]);
                    checkItemSwitch($(trs_form[i]).find('input.field_canadd'), data.forms[i][3]);
                    $(trs_form[i]).find('.field_default').val(data.forms[i][4]);
                    $(trs_form[i]).find('select.field_display').val(data.forms[i][5]).select2();
                    $(trs_form[i]).find('select.field_db_type_form').val(data.forms[i][6]).select2();
                    $(trs_form[i]).find('select.field_form_type_form').val(data.forms[i][7]).select2();
                }

                let detail_table = $("tbody.fields_detail-table");
                detail_table.find("tr").remove();
                let tpl_detail = $("template.fields_detail-tpl").html();
                for (let i = 0; i < data.details.length; i++) {
                    detail_table.append(tpl_detail);
                }

                let trs_detail = detail_table.find("tr");
                for (let i = 0; i < trs_detail.length; i++) {
                    $(trs_detail[i]).find('.field_head').val(data.details[i][0]);
                    $(trs_detail[i]).find('.field_name').val(data.details[i][1]);
                    $(trs_detail[i]).find('select.field_db_type').val(data.details[i][2]).select2();
                }

                toggleItemSwitchOpposite($(".permission.ga_checkbox"), data.permission);
                toggleItemSwitch($(".hide_filter_area.ga_checkbox"), data.hide_filter_area);
                toggleItemSwitch($(".hide_new_button.ga_checkbox"), data.hide_new_button);
                toggleItemSwitch($(".hide_export_button.ga_checkbox"), data.hide_export_button);
                toggleItemSwitch($(".hide_edit_button.ga_checkbox"), data.hide_edit_button);
                toggleItemSwitch($(".hide_pagination.ga_checkbox"), data.hide_pagination);
                toggleItemSwitch($(".hide_delete_button.ga_checkbox"), data.hide_delete_button);
                toggleItemSwitch($(".hide_detail_button.ga_checkbox"), data.hide_detail_button);
                toggleItemSwitch($(".hide_filter_button.ga_checkbox"), data.hide_filter_button);
                toggleItemSwitch($(".hide_row_selector.ga_checkbox"), data.hide_row_selector);
                toggleItemSwitch($(".hide_query_info.ga_checkbox"), data.hide_query_info);
                toggleItemSwitch($(".filter_form_layout.ga_checkbox"), data.filter_form_layout);
                toggleItemSwitch($(".hide_continue_edit_check_box.ga_checkbox"), data.hide_continue_edit_check_box);
                toggleItemSwitch($(".hide_reset_button.ga_checkbox"), data.hide_reset_button);
                toggleItemSwitch($(".hide_continue_new_check_box.ga_checkbox"), data.hide_continue_new_check_box);
                toggleItemSwitch($(".hide_back_button.ga_checkbox"), data.hide_back_button);
            }
            NProgress.done();
        }

        function toggleItemSwitch(obj, val) {
            if (val === "n") {
                $(obj).bootstrapSwitch('state', true);
            } else {
                $(obj).bootstrapSwitch('state', false);
            }
        }

        function toggleItemSwitchOpposite(obj, val) {
            if (val === "n") {
                $(obj).bootstrapSwitch('state', false);
            } else {
                $(obj).bootstrapSwitch('state', true);
            }
        }

        function checkItemSwitch(obj, val) {
            if (val === "y") {
                $(obj).iCheck("check")
            } else {
                $(obj).iCheck("uncheck")
            }
        }

        function getItemSwitchValue(obj) {
            if ($(obj).hasClass("checked")) {
                return "y"
            }
            return "n"
        }

        function getItemObjData() {
            let data = {};
            data.conn = $("select.conn").val();
            data.package = $(".package").val();
            data.pk = $(".pk").val();
            data.path = $(".path").val();
            data.extra_code = $("#extra_code_input").val();
            data.table = $("select.table").val();
            data.table_title = $(".table_title").val();
            data.table_description = $(".table_description").val();
            data.form_title = $(".form_title").val();
            data.form_description = $(".form_description").val();
            data.extra_import_package = $(".extra_import_package").val();
            data.detail_title = $(".detail_title").val();
            data.detail_description = $(".detail_description").val();
            data.detail_display = $("select.detail_display").val();

            let infos = [];
            let trs = $("tbody.fields-table").find("tr");
            for (let i = 0; i < trs.length; i++) {
                infos[i] = [];
                infos[i].push($(trs[i]).find('.field_head').val());
                infos[i].push($(trs[i]).find('.field_name').val());
                infos[i].push(getItemSwitchValue($(trs[i]).find('input.field_filterable').parent()));
                infos[i].push(getItemSwitchValue($(trs[i]).find('input.field_sortable').parent()));
                infos[i].push(getItemSwitchValue($(trs[i]).find('input.field_hide').parent()));
                infos[i].push(getItemSwitchValue($(trs[i]).find('input.info_field_editable').parent()));
                infos[i].push($(trs[i]).find('select.field_db_type').val());
            }
            data.infos = infos;

            let forms = [];
            let trs_form = $("tbody.fields_form-table").find("tr");
            for (let i = 0; i < trs_form.length; i++) {
                forms[i] = [];
                forms[i].push($(trs_form[i]).find('.field_head_form').val());
                forms[i].push($(trs_form[i]).find('.field_name_form').val());
                forms[i].push(getItemSwitchValue($(trs_form[i]).find('input.field_canedit').parent()));
                forms[i].push(getItemSwitchValue($(trs_form[i]).find('input.field_canadd').parent()));
                forms[i].push($(trs_form[i]).find('.field_default').val());
                forms[i].push($(trs_form[i]).find('select.field_display').val());
                forms[i].push($(trs_form[i]).find('select.field_db_type_form').val());
                forms[i].push($(trs_form[i]).find('select.field_form_type_form').val());
            }
            data.forms = forms;

            let details = [];
            let detail_trs = $("tbody.fields_detail-table").find("tr");
            for (let i = 0; i < detail_trs.length; i++) {
                details[i] = [];
                details[i].push($(detail_trs[i]).find('.field_head').val());
                details[i].push($(detail_trs[i]).find('.field_name').val());
                details[i].push($(detail_trs[i]).find('select.field_db_type').val());
            }
            data.details = details;

            data.permission = $("input[name='permission']").val();

            data.hide_filter_area = $("input[name='hide_filter_area']").val();
            data.hide_new_button = $("input[name='hide_new_button']").val();
            data.hide_export_button = $("input[name='hide_export_button']").val();
            data.hide_edit_button = $("input[name='hide_edit_button']").val();
            data.hide_pagination = $("input[name='hide_pagination']").val();
            data.hide_delete_button = $("input[name='hide_delete_button']").val();
            data.hide_detail_button = $("input[name='hide_detail_button']").val();
            data.hide_filter_button = $("input[name='hide_filter_button']").val();
            data.hide_row_selector = $("input[name='hide_row_selector']").val();
            data.hide_query_info = $("input[name='hide_query_info']").val();
            data.filter_form_layout = $("select.filter_form_layout").val();

            data.hide_continue_edit_check_box = $('input[name="hide_continue_edit_check_box"]').val();
            data.hide_reset_button = $('input[name="hide_reset_button"]').val();
            data.hide_continue_new_check_box = $('input[name="hide_continue_new_check_box"]').val();
            data.hide_back_button = $('input[name="hide_back_button"]').val();

            return JSON.stringify(data)
        }

        function conn_req_refresh(selectObj, box, event) {
            $.ajax({
                url: "\/admin\/operation\/_tool_choose_conn",
                type: 'post',
                dataType: 'text',
                data: {
                    'value': $("select.conn").val(),

                    'event': event
                },
                success: function (data) {
                    if (typeof (data) === "string") {
                        data = JSON.parse(data);
                    }
                    if (data.code === 0) {


                        if (selectObj.length > 0) {
                            if (typeof (data.data) === "object") {
                                if (box) {
                                    conn_updateBoxSelections(selectObj, data.data)
                                } else {
                                    if (typeof (selectObj.attr("multiple")) !== "undefined") {
                                        selectObj.html("");
                                    }
                                    selectObj.select2({
                                        data: data.data
                                    });
                                }
                            } else {
                                if (box) {
                                    selectObj.val(data.data).select2()
                                } else {

                                }
                            }
                        } else {
                            $('.table').val(data.data);
                        }


                    } else {
                        swal(data.msg, '', 'error');
                    }
                },
                error: function () {
                    alert('error')
                }
            });
        }
    </script>
    <style>
        .save_table_list {
            position: absolute;
            right: 45px;
            top: 200px;
            background-color: white;
            width: 300px;
            min-height: 50px;
            z-index: 9999;
            display: none;
        }

        .list-group-item.list-head {
            background-color: #5a5a5a;
            border-color: #5a5a5a;
            font-weight: bold;
            color: white;
        }

        .list-group-item.list-group-item-action {
            cursor: pointer;
        }
    </style>
{{end}}`,}