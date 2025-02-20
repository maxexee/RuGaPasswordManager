using System;
using System.ComponentModel.DataAnnotations;

namespace RuGaPasswordManager.DTOs.Users_DTOs.Account
{
    public class LogInDTO
    {
        [Required]
        [EmailAddress]
        public string EmailAddress { get; set; }    =   string.Empty;

        [Required]
        public string Password { get; set; }    =   string.Empty;
    }
}