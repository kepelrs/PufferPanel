<!--
  - Copyright 2019 Padduck, LLC
  -  Licensed under the Apache License, Version 2.0 (the "License");
  -  you may not use this file except in compliance with the License.
  -  You may obtain a copy of the License at
  -          http://www.apache.org/licenses/LICENSE-2.0
  -  Unless required by applicable law or agreed to in writing, software
  -  distributed under the License is distributed on an "AS IS" BASIS,
  -  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  -  See the License for the specific language governing permissions and
  -  limitations under the License.
  -->

<template>
  <v-container>
    <h1 v-text="$t('servers.Add')" />
    <v-stepper v-model="currentStep">
      <v-stepper-header>
        <v-stepper-step step="1" :complete="currentStep > 1" />
        <v-divider />
        <v-stepper-step step="2" :complete="currentStep > 2" />
        <v-divider />
        <v-stepper-step step="3" :complete="currentStep > 3" />
        <v-divider />
        <v-stepper-step step="4" :complete="currentStep > 4" />
      </v-stepper-header>
      <v-stepper-items>
        <v-stepper-content step="1">
          <h3 v-text="$t('servers.SelectTemplate')" />
          <v-text-field v-model="templateFilter" :placeholder="$t('common.Search')" />
          <v-expansion-panels>
            <v-expansion-panel v-for="template in filterTemplates(templates, templateFilter)" :key="template.name" @click="loadTemplateData(template.name)">
              <v-expansion-panel-header v-text="template.display" />
              <v-expansion-panel-content>
                <v-row v-if="templateData[template.name] === undefined">
                  <v-col cols="5" />
                  <v-col cols="2">
                    <v-progress-circular
                      indeterminate
                      class="mr-2"
                    />
                    <strong v-text="$t('common.Loading')" />
                  </v-col>
                </v-row>
                <span v-else v-html="markdown(templateData[template.name].readme || $t('servers.NoReadme'))" />
                <br />
                <v-btn color="primary" @click="selectTemplate(template.name)" v-text="$t('servers.SelectThisTemplate')" large block />
              </v-expansion-panel-content>
            </v-expansion-panel>
          </v-expansion-panels>
        </v-stepper-content>

        <v-stepper-content step="2">
          <v-row>
            <v-col cols="12">
              <h3 v-text="$t('servers.Name')" class="mb-4" />
              <v-text-field
                id="nameInput"
                v-model="serverName"
                outlined
              />
            </v-col>
          </v-row>

          <v-row>
            <v-col cols="12">
              <h3 v-text="$t('nodes.Node')" class="mb-4" />
              <v-select
                id="nodeSelect"
                v-model="selectedNode"
                outlined
                :disabled="loadingNodes"
                :items="nodes"
                single-line
                :no-data-text="$t('errors.ErrNoNodes')"
                :placeholder="$t('servers.SelectNode')"
              />
            </v-col>
          </v-row>

          <v-row>
            <v-col cols="12">
              <h3 v-text="$t('servers.Environment')" class="mb-4" />
              <v-select
                id="environmentSelect"
                v-model="selectedEnvironment"
                :disabled="loadingTemplates"
                :items="environments"
                outlined
                :placeholder="$t('servers.SelectEnvironment')"
              />
              <div v-if="selectedEnvironment && environments[selectedEnvironment]">
                <div v-for="(val, key) in environments[selectedEnvironment].metadata">
                  <v-text-field
                    v-model="environments[selectedEnvironment].metadata[key]"
                    outlined
                    :label="$t('env.' + environments[selectedEnvironment].type + '.' + key)"
                  />
                </div>
              </div>
            </v-col>
          </v-row>

          <v-row>
            <v-col>
              <v-btn
                large
                block
                color="error"
                @click="currentStep = 1"
                v-text="$t('common.Back')"
              />
            </v-col>
            <v-col>
              <v-btn
                large
                block
                color="primary"
                :disabled="selectedNode == null || selectedEnvironment == null || serverName == ''"
                @click="currentStep = 3"
                v-text="$t('common.Next')"
              />
            </v-col>
          </v-row>
        </v-stepper-content>

        <v-stepper-content step="3">
          <v-row>
            <v-col cols="12">
              <h3 v-text="$t('users.Users')" class="mb-4" />
              <v-text-field
                v-model="userInput"
                outlined
                :placeholder="$t('servers.TypeUsername')"
              />
              <v-list v-if="users.length > 0 || selectedUsers.length > 0">
                <v-subheader
                  v-if="users.length > 0"
                  v-text="$t('users.Add')"
                />
                <v-list-item-group v-if="users.length > 0">
                  <v-list-item
                    v-for="user in users"
                    :key="user.value"
                    @click="selectUser(user.value)"
                  >
                    <v-list-item-icon>
                      <v-icon>mdi-plus</v-icon>
                    </v-list-item-icon>
                    <v-list-item-content>
                      <v-list-item-title v-text="user.text" />
                    </v-list-item-content>
                  </v-list-item>
                </v-list-item-group>
                <v-subheader
                  v-if="selectedUsers.length > 0"
                  v-text="$t('users.Users')"
                />
                <v-list-item-group v-if="selectedUsers.length > 0">
                  <v-list-item
                    v-for="user in selectedUsers"
                    @click="removeUser(user)"
                  >
                    <v-list-item-icon>
                      <v-icon>mdi-minus</v-icon>
                    </v-list-item-icon>
                    <v-list-item-content>
                      <v-list-item-title v-text="user" />
                    </v-list-item-content>
                  </v-list-item>
                </v-list-item-group>
              </v-list>
            </v-col>
          </v-row>

          <v-row>
            <v-col>
              <v-btn
                large
                block
                color="error"
                @click="currentStep = 2"
                v-text="$t('common.Back')"
              />
            </v-col>
            <v-col>
              <v-btn
                large
                block
                color="primary"
                :disabled="selectedUsers.length === 0"
                @click="currentStep = 4"
                v-text="$t('common.Next')"
              />
            </v-col>
          </v-row>
        </v-stepper-content>

        <v-stepper-content step="4">
          <v-row v-if="Object.keys(formData).length > 0">
            <v-col cols="12">
              <v-card-title v-text="$t('common.Options')" />
              <v-row>
                <v-col
                  v-for="item in formData"
                  v-if="!item.internal"
                  cols="12"
                >
                  <v-text-field
                    v-if="item.type === 'integer'"
                    v-model="item.value"
                    type="number"
                    :required="item.required"
                    :hint="item.desc"
                    persistent-hint
                    :label="item.display"
                    outlined
                  >
                    <template slot="message"><div v-html="item.desc" /></template>
                  </v-text-field>
                  <v-switch
                    v-else-if="item.type === 'boolean'"
                    v-model="item.value"
                    class="mt-0 mb-4"
                    :required="item.required"
                    :hint="item.desc"
                    persistent-hint
                    :label="item.display"
                  >
                    <template slot="message"><div v-html="item.desc" /></template>
                  </v-switch>
                  <v-select
                    v-else-if="item.type === 'option'"
                    v-model="item.value"
                    :items="item.options.map(function (option) { return { value: option.value, text: option.display }})"
                    :hint="item.desc"
                    persistent-hint
                    :label="item.display"
                    outlined
                  >
                    <template slot="message"><div v-html="item.desc" /></template>
                  </v-select>
                  <v-text-field
                    v-else
                    v-model="item.value"
                    :required="item.required"
                    :hint="item.desc"
                    persistent-hint
                    :label="item.display"
                    outlined
                  >
                    <template slot="message"><div v-html="item.desc" /></template>
                  </v-text-field>
                </v-col>
              </v-row>
            </v-col>
          </v-row>

          <v-row>
            <v-col>
              <v-btn
                large
                block
                color="error"
                @click="currentStep = 3"
                v-text="$t('common.Back')"
              />
            </v-col>
            <v-col>
              <v-btn
                large
                block
                color="primary"
                :disabled="!canCreate"
                @click="submitCreate"
                v-text="$t('common.Create')"
              />
            </v-col>
          </v-row>
        </v-stepper-content>
      </v-stepper-items>
    </v-stepper>
  </v-container>
</template>

<script>
import axios from 'axios'
import { handleError } from '@/utils/api'
import markdown from '@/utils/markdown'

const CancelToken = axios.CancelToken

export default {
  data () {
    return {
      nodes: [],
      selectedNode: null,
      templateFilter: '',
      templates: [],
      templateData: {},
      selectedTemplate: '',
      formData: {},

      loadingNodes: true,
      loadingTemplates: true,

      searchingUsers: true,
      users: [],
      selectedUser: null,
      selectedUsers: [],
      userInput: null,
      userCancelSearch: CancelToken.source(),

      serverName: '',

      selectedEnvironment: null,
      environments: [],

      currentStep: 1
    }
  },
  computed: {
    canCreate () {
      if (this.loadingTemplates || this.loadingNodes) {
        return false
      }

      if (!this.selectedTemplate || this.selectedTemplate === '') {
        return false
      }

      if (this.selectedUsers.length === 0) {
        return false
      }

      if (!this.selectedEnvironment) {
        return false
      }

      for (const k in this.templateData[this.selectedTemplate].data) {
        const data = this.templateData[this.selectedTemplate].data[k]
        if (data.type === 'boolean') {
          continue
        }
        if (data.required && !data.value) {
          return false
        }
      }

      return true
    }
  },
  watch: {
    selectedTemplate (newVal) {
      if (!newVal || newVal === '') {
        return
      }
      this.formData = this.templateData[newVal].data
      this.environments = []
      for (const k in this.templateData[newVal].supportedEnvironments) {
        const env = this.templateData[newVal].supportedEnvironments[k]
        this.environments.push({
          value: k,
          text: this.$t('env.' + env.type + '.name'),
          metadata: env.metadata,
          type: env.type
        })
      }

      const env = this.templateData[newVal].environment
      let def = null
      if (env && env.type) {
        def = env.type
      }

      if (def) {
        for (const k in this.environments) {
          if (this.environments[k].type === def) {
            this.selectedEnvironment = k
            break
          }
        }
      } else {
        this.selectedEnvironment = null
      }

      for (const key in this.formData) {
        if (this.formData[key].type === 'boolean') {
          this.formData[key].value = this.formData[key].value === 'true'
        }
      }
    },
    userInput (newVal) {
      if (!newVal || newVal === '') {
        this.users = []
      } else {
        new Promise((resolve, reject) => {
          this.findUsers(newVal)
          resolve()
        })
      }
    }
  },
  mounted () {
    this.nodes = [{
      value: null,
      disabled: true,
      text: this.$t('common.Loading')
    }]
    this.getTemplates()
    this.getNodes()
  },
  methods: {
    getTemplates () {
      const ctx = this
      this.loadingTemplates = true
      this.templates = [{
        value: null,
        disabled: true,
        text: this.$t('common.Loading')
      }]
      this.templateData = {}
      this.selectedTemplate = null
      this.$http.get('/api/templates').then(response => {
        ctx.templates = response.data
        ctx.loadingTemplates = false
      }).catch(handleError(ctx))
    },
    getNodes () {
      const ctx = this
      this.$http.get('/api/nodes').then(response => {
        ctx.nodes = []
        for (let i = 0; i < response.data.length; i++) {
          const node = response.data[i]
          ctx.nodes.push({
            value: node.id,
            text: node.name
          })
        }

        if (ctx.nodes.length === 1) {
          ctx.selectedNode = ctx.nodes[0].value
        }

        ctx.loadingNodes = false
      }).catch(handleError(ctx))
    },
    findUsers () {
      const ctx = this
      this.searchingUsers = true
      this.userCancelSearch.cancel()
      this.userCancelSearch = CancelToken.source()
      this.$http.get(`/api/users?username=${this.userInput}*`, {
        cancelToken: this.userCancelSearch.token
      }).then(response => {
        ctx.users = []
        for (let i = 0; i < Math.min(response.data.users.length, 5); i++) {
          const user = response.data.users[i]
          ctx.users.push({
            value: user.username,
            text: user.username + ' <' + user.email + '>'
          })
        }
        ctx.searchingUsers = false
        ctx.users.sort()
      }).catch(handleError(ctx))
    },
    submitCreate () {
      const ctx = this
      const data = this.templateData[this.selectedTemplate]
      for (const item in data.data) {
        switch (data.data[item].type) {
          case 'integer':
            data.data[item].value = Number(data.data[item].value)
            break
          case 'boolean':
            data.data[item].value = Boolean(data.data[item].value)
            break
        }
      }
      data.node = this.selectedNode
      data.users = this.selectedUsers
      data.name = this.serverName !== '' ? this.serverName : undefined
      data.environment = {
        type: this.environments[this.selectedEnvironment].type,
        metadata: this.environments[this.selectedEnvironment].metadata
      }
      this.$http.post('/api/servers', data).then(response => {
        ctx.$router.push({ name: 'Server', params: { id: response.data.id } })
      }).catch(handleError(ctx))
    },
    selectUser (username) {
      if (!username || username === '') {
        return
      }
      for (let i = 0; i < this.selectedUsers.length; i++) {
        if (this.selectedUsers[i] === username) {
          return
        }
      }
      this.userInput = null
      this.selectedUsers.push(username)
      this.selectedUsers.sort()
      this.selectedUser = null
      this.users = []
    },
    removeUser (username) {
      for (let i = 0; i < this.selectedUsers.length; i++) {
        if (this.selectedUsers[i] === username) {
          this.selectedUsers.splice(i, 1)
          return
        }
      }
    },
    loadTemplateData (template) {
      if (!template) return
      if (!this.templateData[template]) {
        const ctx = this
        this.$http.get(`/api/templates/${template}`).then(response => {
          if (response.status >= 200 && response.status < 300) {
            ctx.templateData[template] = response.data
            // recreate object to make vues refernce equality check fail and rerender neccessary components
            ctx.templateData = { ...ctx.templateData }
          }
        }).catch(handleError(ctx))
      }
    },
    selectTemplate (template) {
      this.loadTemplateData(template)
      this.selectedTemplate = template
      this.currentStep = 2
    },
    filterTemplates (templates, filter) {
      return templates.filter(t => {
        if (filter.trim() === '') {
          return true
        } else {
          return t.display.toLowerCase().indexOf(filter.trim().toLowerCase()) > -1
        }
      })
    },
    markdown
  }
}
</script>
