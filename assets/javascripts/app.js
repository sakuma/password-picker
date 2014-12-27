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
        key: null,
        body: [
          {
            key: null,
            value: null
          },
          {
            key: null,
            value: null
          }
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
            // console.log(res);
            // console.log(res.text);
            // console.log("-----")
            if (res.status === 200) {
              var password = JSON.parse(res.text);
              password.Body = JSON.parse(password.Body);
              console.log(password);
              self.passwords.unshift(password);
            } else {
              console.log("faile");
            }
          });
        return false;
      }
    }
    // methods: {
    //     addUser: function (e) {
    //         e.preventDefault()
    //         console.log(this)
    //         if (this.validation.name && this.validation.email) {
    //             Users.push(this.newUser)
    //             this.newUser = {}
    //         }
    //     },
    //     removeUser: function (user) {
    //         new Firebase(baseURL + 'users/' + user.id).remove()
    //     }
    // }
});
