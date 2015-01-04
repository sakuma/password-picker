'use strict';

new Vue({
    el: '#app',
    created: function () {
      this.fetchData();
    },
    data: {
      passwords: [],
      newPassword: {
        title: null,
        note: null,
        attributes: [
          { key: null, value: null },
        ]
      }
    },
    methods: {
      fetchData: function () {
        var self = this;
        window.superagent.get('/passwords').send().end(function (res) {
          var tmp = JSON.parse(res.text);
          for (var i=0; i < tmp.length; ++i) {
            tmp[i].EditMode = false;
          }
          self.passwords = tmp;
        });
      },

      addPassword: function () {
        var self = this;
        window.superagent.post('/passwords')
          .send({
            title: this.newPassword.title,
            attributes: this.newPassword.attributes,
            note: this.newPassword.note
          })
          .end(function (res) {
            if (res.status === 200) {
              console.log(res)
              var password = JSON.parse(res.text);
              self.passwords.$set(0, password);
            } else {
              console.log("failure");
              console.log(res.text);
            }
          });
        return false;
      },

      deletePassword: function (password, index) {
        var vm = this;
        window.superagent.del('/passwords/' + password.Id).send({})
          .end(function (res) {
            if (res.status === 200) {
              vm.passwords.$remove(index);
            } else {
              // TODO: toast
              console.log("failure");
              console.log(res.text);
            }
          });
        return false;
      },

      addAttribute: function () {
        this.newPassword.attributes.$add({ key: null, value: null });
      },

      delAttribute: function () {
        this.newPassword.attributes.pop();
      },

      toggleEditMode: function (index) {
        this.passwords[index].EditMode = !this.passwords[index].EditMode;
      }
    }
});

Vue.component('password', {
  template: '#password-template',
  methods: {
    updatePassword: function (obj) {
      var self = obj;
      window.superagent.put('/passwords/' + self.Id)
      .send({
        title: self.Title,
        attributes: self.Attributes,
        note: self.Note
      })
      .end(function (res) {
        if (res.status === 200) {
          console.log("success")
          var password = JSON.parse(res.text);
          self.EditMode = false;
        } else {
          console.log("failure");
          console.log(res.text);
        }
      });
    }
  }
});

$(document).ready(function(){
  $('.modal-trigger').leanModal();
});
