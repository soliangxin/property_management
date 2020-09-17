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
      <el-table-column align="center" label="设备编号" width="150">
        <template slot-scope="scope">
          {{ scope.row.code }}
        </template>
      </el-table-column>
      <el-table-column label="设备名称" width="150" align="center">
        <template slot-scope="scope">
          {{ scope.row.name }}
        </template>
      </el-table-column>
      <el-table-column label="品牌" width="110" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.brand }}</span>
        </template>
      </el-table-column>
      <el-table-column label="所属小区编号" width="160" align="center">
        <template slot-scope="scope">
          {{ scope.row.community_code }}
        </template>
      </el-table-column>
      <el-table-column label="购买价格" width="110" align="center">
        <template slot-scope="scope">
          {{ scope.row.price }}
        </template>
      </el-table-column>
      <el-table-column label="购买数量" width="110" align="center">
        <template slot-scope="scope">
          {{ scope.row.quantity }}
        </template>
      </el-table-column>
      <el-table-column label="预计使用年限" width="110" align="center">
        <template slot-scope="scope">
          {{ scope.row.durable_years }}
        </template>
      </el-table-column>
      <el-table-column align="center" prop="created_at" label="购买时间" width="210">
        <template slot-scope="scope">
          <i class="el-icon-time" />
          <span>{{ scope.row.buy_time }}</span>
        </template>
      </el-table-column>
      <el-table-column label="备注" align="center">
        <template slot-scope="scope">
          {{ scope.row.remark }}
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script>
import { getAssetList } from '@/api/asset'

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
      getAssetList().then(response => {
        this.list = response.data.items
        this.listLoading = false
      })
    }
  }
}
</script>
