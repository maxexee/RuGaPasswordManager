using System;

namespace RuGaPasswordManager.DTOs.Users_DTOs.Account
{
    public class RetunNewUser
    {
        public string UserName { get; set; }    =   string.Empty;
        public string Email { get; set; }    =   string.Empty;
        public string Token { get; set; }    =   string.Empty;
    }
}