using System;
using System.ComponentModel.DataAnnotations;

namespace RuGaPasswordManager.DTOs.Users_DTOs.Account
{
    public class RegisterDTO
    {
        [Required]
        public string UserName { get; set; }    =   string.Empty;

        [Required]
        [EmailAddress]
        public string EmailAddress { get; set; }    =   string.Empty;

        [Required]
        public string Password { get; set; }    =   string.Empty;
    }
}