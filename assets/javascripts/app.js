'use strict';

new Vue({
    el: '#app',
    created: function () {
      this.fetchData();
    },
    data: {
      visibilityForm: false,
      passwords: [],
      newPassword: {
        title: null,
        note: null,
        body: [
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
            tmp[i].Body = JSON.parse(tmp[i].Body);
            tmp[i].EditMode = false;
          }
          self.passwords = tmp;
        });
      },

      addPassword: function () {
        var self = this;
        window.superagent.post('/passwords')
          .send({
            title: self.newPassword.title,
            body: JSON.stringify(this.newPassword.body)
          })
          .end(function (res) {
            if (res.status === 200) {
              var password = JSON.parse(res.text);
              password.Body = JSON.parse(password.Body);
              self.passwords.unshift(password);
              self.visibilityForm = false;
            } else {
              console.log("failure");
              console.log(res.text);
            }
          });
        return false;
      },

      visibleAddPassword: function () {
        this.visibilityForm = true;
      },

      addAttribute: function () {
        this.newPassword.body.push({ key: null, value: null });
      },

      delAttribute: function () {
        this.newPassword.body.pop();
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
        body: JSON.stringify(self.Body),
        note: self.Note
      })
      .end(function (res) {
        if (res.status === 200) {
          console.log("success")
          var password = JSON.parse(res.text);
          password.Body = JSON.parse(password.Body);
          self.EditMode = false;
        } else {
          console.log("failure");
          console.log(res.text);
        }
      });
    }
  }
});
