<template>
  <BasicLayout>
    <template #wrapper>
      <el-card class="box-card">
        <el-form
          ref="queryForm"
          :model="queryParams"
          :inline="true"
          label-width="68px"
        >
          <el-form-item label="姓名" prop="interviewId">
            <el-input
              v-model="queryParams.interviewName"
              placeholder="请输入候选人姓名"
              clearable
              size="small"
              @keyup.enter.native="handleQuery"
            />
          </el-form-item>
          <el-form-item label="岗位" prop="postName">
            <el-select
              v-model="queryParams.postName"
              placeholder="岗位名称"
              clearable
              size="small"
            >
              <el-option
                v-for="dict in postOptions"
                :key="dict.postCode"
                :label="dict.postName"
                :value="dict.postCode"
              />
            </el-select>
          </el-form-item>
          <el-form-item label="状态" prop="status">
            <el-select
              v-model="queryParams.status"
              placeholder="候选人状态"
              clearable
              size="small"
            >
              <el-option
                v-for="dict in statusOptions"
                :key="dict.value"
                :label="dict.label"
                :value="dict.value"
              />
            </el-select>
          </el-form-item>
          <el-form-item label="创建时间">
            <el-date-picker
              v-model="dateRange"
              size="small"
              style="width: 240px"
              value-format="yyyy-MM-dd"
              type="daterange"
              range-separator="-"
              start-placeholder="开始日期"
              end-placeholder="结束日期"
            />
          </el-form-item>
          <el-form-item>
            <el-button
              type="primary"
              icon="el-icon-search"
              size="mini"
              @click="handleQuery"
            >搜索</el-button>
            <el-button
              icon="el-icon-refresh"
              size="mini"
              @click="resetQuery"
            >重置</el-button>
          </el-form-item>
        </el-form>

        <el-row :gutter="10" class="mb8">
          <el-col :span="1.5">
            <el-button
              v-permisaction="['admin:interview:add']"
              type="primary"
              icon="el-icon-plus"
              size="mini"
              @click="handleAdd"
            >新增</el-button>
          </el-col>
          <el-col :span="1.5">
            <el-button
              v-permisaction="['admin:interview:edit']"
              type="success"
              icon="el-icon-edit"
              size="mini"
              :disabled="single"
              @click="handleUpdate"
            >修改</el-button>
          </el-col>
          <el-col :span="1.5">
            <el-button
              v-permisaction="['admin:interview:remove']"
              type="danger"
              icon="el-icon-delete"
              size="mini"
              :disabled="multiple"
              @click="handleDelete"
            >删除</el-button>
          </el-col>
          <el-col :span="1.5">
            <el-button
              v-permisaction="['admin:interview:export']"
              type="warning"
              icon="el-icon-download"
              size="mini"
              @click="handleExport"
            >导出</el-button>
          </el-col>
        </el-row>

        <el-table
          v-loading="loading"
          :data="postList"
          border
          @selection-change="handleSelectionChange"
        >
          <el-table-column type="selection" width="55" align="center" />
          <el-table-column
            label="编号"
            width="80"
            sortable="custom"
            align="center"
            prop="interviewId"
          />
          <el-table-column
            label="候选人姓名"
            width="120"
            align="center"
            prop="interviewName"
          />
          <!-- <el-table-column
            label="岗位名称"
            width="120"
            align="center"
            :show-overflow-tooltip="true"
            prop="postCode"
          /> -->
          <el-table-column
            label="岗位名称"
            align="center"
            width="120"
            prop="postCode"
            :show-overflow-tooltip="true"
            :formatter="postFormat"
          >
            <template slot-scope="scope">
              <el-tag>{{ postFormat(scope.row) }}</el-tag>
            </template>
          </el-table-column>
          <el-table-column
            label="状态"
            align="center"
            prop="status"
            width="120"
            :formatter="statusFormat"
          >
            <template slot-scope="scope">
              <el-tag
                :type="scope.row.status === 3 ? 'danger' : 'success'"
                disable-transitions
              >{{ statusFormat(scope.row) }}</el-tag>
            </template>
          </el-table-column>
          <el-table-column label="备注" align="center" prop="remark" :show-overflow-tooltip="true">
            <template slot-scope="scope">
              <span>{{ scope.row.remark }}</span>
            </template>
          </el-table-column>
          <el-table-column
            label="创建时间"
            align="center"
            prop="createdAt"
            width="180"
          >
            <template slot-scope="scope">
              <span>{{ parseTime(scope.row.createdAt) }}</span>
            </template>
          </el-table-column>
          <el-table-column
            label="操作"
            align="center"
            width="120"
            class-name="small-padding fixed-width"
          >
            <template slot-scope="scope">
              <el-button
                v-permisaction="['admin:interview:edit']"
                size="mini"
                type="text"
                icon="el-icon-edit"
                @click="handleUpdate(scope.row)"
              >修改</el-button>
              <el-button
                v-permisaction="['admin:interview:remove']"
                size="mini"
                type="text"
                icon="el-icon-delete"
                @click="handleDelete(scope.row)"
              >删除</el-button>
            </template>
          </el-table-column>
        </el-table>

        <pagination
          v-show="total > 0"
          :total="total"
          :page.sync="queryParams.pageIndex"
          :limit.sync="queryParams.pageSize"
          @pagination="getList"
        />

        <!-- 添加或修改岗位对话框 -->
        <el-dialog
          :title="title"
          :visible.sync="open"
          width="500px"
          :close-on-click-modal="false"
        >
          <el-form ref="form" :model="form" :rules="rules" label-width="80px">
            <el-form-item label="岗位名称" prop="postId">
              <el-select
                v-model="form.postId"
                placeholder="岗位名称"
                clearable
                size="small"
              >
                <el-option
                  v-for="dict in postOptions"
                  :key="dict.postId"
                  :label="dict.postName"
                  :value="dict.postId"
                />
              </el-select>
            </el-form-item>
            <el-form-item label="候选人姓名" prop="interviewName">
              <el-input v-model="form.interviewName" placeholder="请输入候选人姓名" />
            </el-form-item>
            <el-form-item label="岗位状态" prop="status">
              <el-radio-group v-model="form.status">
                <el-radio
                  v-for="dict in statusOptions"
                  :key="dict.value"
                  :label="dict.value"
                >{{ dict.label }}</el-radio>
              </el-radio-group>
            </el-form-item>
            <el-form-item label="备注" prop="remark">
              <el-input
                v-model="form.remark"
                type="textarea"
                placeholder="请输入内容"
              />
            </el-form-item>
          </el-form>
          <div slot="footer" class="dialog-footer">
            <el-button type="primary" @click="submitForm">确 定</el-button>
            <el-button @click="cancel">取 消</el-button>
          </div>
        </el-dialog>
      </el-card>
    </template>
  </BasicLayout>
</template>

<script>
import {
  listPost
} from '@/api/admin/sys-post'
import {
  listInterview,
  getInterview,
  delInterview,
  addInterview,
  updateInterview
} from '@/api/admin/sys-interview'
import { formatJson } from '@/utils'

export default {
  name: 'SysPostManage',
  data() {
    return {
      // 遮罩层
      loading: true,
      // 选中数组
      ids: [],
      // 非单个禁用
      single: true,
      // 非多个禁用
      multiple: true,
      // 总条数
      total: 0,
      // 岗位表格数据
      postList: [],
      // 弹出层标题
      title: '',
      // 是否显示弹出层
      open: false,
      // 状态数据字典
      statusOptions: [],
      // 岗位数据字典
      postOptions: [],
      postMap: {},
      // 查询参数
      queryParams: {
        pageIndex: 1,
        pageSize: 10,
        interviewName: undefined,
        postCode: undefined,
        postName: undefined,
        status: undefined
      },
      // 表单参数
      form: {},
      // 表单校验
      rules: {
        postName: [
          { required: true, message: '岗位名称不能为空', trigger: 'blur' }
        ],
        interviewName: [
          { required: true, message: '候选人姓名不能为空', trigger: 'blur' }
        ]
      }
    }
  },
  created() {
    this.getPostList()
    this.getDicts('sys_interview_disable').then((response) => {
      this.statusOptions = response.data
    })
  },
  methods: {
    /** 查询岗位列表 */
    getPostList() {
      this.loading = true
      listPost({
        pageIndex: 1,
        pageSize: 10000
      }).then((response) => {
        response.data.list.forEach((element) => {
          const { postId, postName } = element
          this.postMap[postId] = postName
        })
        this.postOptions = response.data.list

        this.getList()
      })
    },
    /** 查询候选人列表 */
    getList() {
      this.loading = true
      listInterview(this.queryParams).then((response) => {
        this.postList = response.data.list
        this.total = response.data.count
        this.loading = false
      })
    },
    // 岗位状态字典翻译
    statusFormat(row) {
      return this.selectDictLabel(this.statusOptions, row.status)
    },
    // 岗位状态字典翻译
    postFormat(row) {
      return this.postMap[row.postCode]
    },
    // 取消按钮
    cancel() {
      this.open = false
      this.reset()
    },
    // 表单重置
    reset() {
      this.form = {
        interviewId: undefined,
        postCode: undefined,
        interviewName: undefined,
        status: '1',
        remark: undefined
      }
      this.resetForm('form')
    },
    /** 搜索按钮操作 */
    handleQuery() {
      this.queryParams.pageIndex = 1
      this.getList()
    },
    /** 重置按钮操作 */
    resetQuery() {
      this.resetForm('queryForm')
      this.handleQuery()
    },
    // 多选框选中数据
    handleSelectionChange(selection) {
      this.ids = selection.map((item) => item.interviewId)
      this.single = selection.length !== 1
      this.multiple = !selection.length
    },
    /** 新增按钮操作 */
    handleAdd() {
      this.reset()
      this.open = true
      this.title = '添加候选人'
    },
    /** 修改按钮操作 */
    handleUpdate(row) {
      this.reset()

      const interviewId = row.interviewId || this.ids
      getInterview(interviewId).then((response) => {
        this.form = response.data
        this.form.status = String(this.form.status)
        this.open = true
        this.title = '修改候选人'
      })
    },
    /** 提交按钮 */
    submitForm: function() {
      this.$refs['form'].validate((valid) => {
        if (valid) {
          this.form.status = parseInt(this.form.status)
          if (this.form.interviewId !== undefined) {
            updateInterview(this.form, this.form.interviewId).then((response) => {
              if (response.code === 200) {
                this.msgSuccess(response.msg)
                this.open = false
                this.getList()
              } else {
                this.msgError(response.msg)
              }
            })
          } else {
            addInterview(this.form).then((response) => {
              if (response.code === 200) {
                this.msgSuccess(response.msg)
                this.open = false
                this.getList()
              } else {
                this.msgError(response.msg)
              }
            })
          }
        }
      })
    },
    /** 删除按钮操作 */
    handleDelete(row) {
      // const interviewIds = row.interviewId || this.ids
      const Ids = (row.interviewId && [row.interviewId]) || this.ids
      this.$confirm('是否确认删除候选人编号为"' + Ids + '"的数据项?', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
        .then(function() {
          return delInterview({ ids: Ids })
        })
        .then((response) => {
          if (response.code === 200) {
            this.msgSuccess(response.msg)
            this.open = false
            this.getList()
          } else {
            this.msgError(response.msg)
          }
        })
        .catch(function() {})
    },
    /** 导出按钮操作 */
    handleExport() {
      // const queryParams = this.queryParams
      this.$confirm('是否确认导出所有岗位数据项?', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
        .then(() => {
          this.downloadLoading = true
          import('@/vendor/Export2Excel').then((excel) => {
            const tHeader = [
              '候选人编号',
              '候选人姓名',
              '岗位名称',
              '排序',
              '创建时间'
            ]
            const filterVal = [
              'interviewId',
              'postCode',
              'postName',
              'sort',
              'createdAt'
            ]
            const list = this.postList
            const data = formatJson(filterVal, list)
            excel.export_json_to_excel({
              header: tHeader,
              data,
              filename: '岗位管理',
              autoWidth: true, // Optional
              bookType: 'xlsx' // Optional
            })
            this.downloadLoading = false
          })
        })
        .catch(function() {})
    }
  }
}
</script>
