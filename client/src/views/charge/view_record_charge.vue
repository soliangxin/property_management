<template>
  <div class="app-container">
    <el-table
      v-loading="listLoading"
      :data="list"
      element-loading-text="Loading"
      border
      fit
      highlight-current-row
    >
      <el-table-column align="center" label="缴费项目编号" width="160">
        <template slot-scope="scope">
          {{ scope.row.project_code }}
        </template>
      </el-table-column>
      <el-table-column label="房产编号" width="170" align="center">
        <template slot-scope="scope">
          {{ scope.row.house_code }}
        </template>
      </el-table-column>
      <el-table-column label="应收金额" width="110" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.amount_total }}</span>
        </template>
      </el-table-column>
      <el-table-column label="实收金额" width="110" align="center">
        <template slot-scope="scope">
          {{ scope.row.amount_paid }}
        </template>
      </el-table-column>
      <el-table-column label="所属小区编号" width="180" align="center">
        <template slot-scope="scope">
          {{ scope.row.community_code }}
        </template>
      </el-table-column>
      <el-table-column align="center" prop="created_at" label="缴费时间" width="210">
        <template slot-scope="scope">
          <i class="el-icon-time" />
          <span>{{ scope.row.pay_time }}</span>
        </template>
      </el-table-column>
      <el-table-column label="备注">
        <template slot-scope="scope">
          {{ scope.row.remark }}
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script>
import { getChangeRecordList } from '@/api/charge'

export default {
  filters: {
    statusFilter(status) {
      const statusMap = {
        published: 'success',
        draft: 'gray',
        deleted: 'danger'
      }
      return statusMap[status]
    }
  },
  data() {
    return {
      list: null,
      listLoading: true
    }
  },
  created() {
    this.fetchData()
  },
  methods: {
    fetchData() {
      this.listLoading = true
      getChangeRecordList().then(response => {
        this.list = response.data.items
        this.listLoading = false
      })
    }
  }
}
</script>
