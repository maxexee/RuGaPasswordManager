using System;
using System.Diagnostics;
using Microsoft.AspNetCore.Mvc;
using Microsoft.AspNetCore.Identity;
using Microsoft.EntityFrameworkCore;

using RuGa.Data.Identities;
using RuGaPasswordManager.DTOs.Users_DTOs.Account;
using RuGaPasswordManager.Interfaces.Account_Interfaces;

namespace RuGaPasswordManager.Controllers.Users.Account
{
    [ApiController]
    [Route("api/[controller]")]
    public class AccountGeneralActions : ControllerBase
    {
        private readonly    ITokenService _tokenService;
        private readonly    UserManager<AppUser> _userManager;
        private readonly    SignInManager<AppUser> _signInManager;
        public AccountGeneralActions(UserManager<AppUser>   userManager,    SignInManager<AppUser>  signInManager,  ITokenService   tokenService)
        {
            _userManager    =   userManager;
            _signInManager  =   signInManager;
            _tokenService   =   tokenService;
        }

        [HttpPost("registerUser")]
        public async Task<IActionResult> RegisterNewUser([FromBody] RegisterDTO    registerDTO){
            try
            {
                if(!ModelState.IsValid) return  BadRequest(ModelState);

                var newUserObj =   new AppUser {
                    UserName    =   registerDTO.UserName,
                    Email       =   registerDTO.EmailAddress,
                };

                var newUserCreation =   await   _userManager.CreateAsync(newUserObj,    registerDTO.Password);
                if (newUserCreation.Succeeded)
                {
                    var roleResult  =   await   _userManager.AddToRoleAsync(newUserObj,    "User");
                    if (roleResult.Succeeded)
                    {
                        return  Ok(
                            new  RetunNewUser    {
                                UserName    =   newUserObj.UserName,
                                Email       =   newUserObj.Email,
                                Token       =   _tokenService.CreateToken(newUserObj)
                            }
                        );
                    }else{
                        return BadRequest(roleResult.Errors);
                    }
                }else{
                    return  BadRequest(newUserCreation.Errors);
                }
            }
            catch (Exception    ex)
            {
                // Get stack trace for the exception with source file information
                var st = new StackTrace(ex, true);
                // Get the top stack frame
                var frame = st.GetFrame(0);
                // Get the line number from the stack frame
                var line = frame?.GetFileLineNumber();
                return  StatusCode(500, ex.Message);
            }
        }

        [HttpPost("loginUser")]
        public  async   Task<IActionResult> LoginUser(LogInDTO  logInDTO){
            try
            {
                if (!ModelState.IsValid)    return  StatusCode(5000);

                // ENCONTRAR EL USUARIO.
                var user = await _userManager.Users.FirstOrDefaultAsync(x   =>  x.Email == logInDTO.EmailAddress);
                if (user    ==  null)   return  Unauthorized("Invalid Username");

                // VERIFICAR CONTRASEÑA.
                var result  =   await   _signInManager.CheckPasswordSignInAsync(user,   logInDTO.Password,  false);
                if (!result.Succeeded)  return  Unauthorized("User not found and/or passwordIncorrect");


                // RETORNO DEL USUARIO Y ACCESO.
                Console.WriteLine($"{user.UserName} Usuario valido.");
                return  Ok(
                    new RetunNewUser    {
                        UserName=   user.UserName,
                        Email   =   user.Email,
                        Token   =   _tokenService.CreateToken(user)
                    }
                );
            }
            catch (Exception    ex)
            {
                // Get stack trace for the exception with source file information
                var st = new StackTrace(ex, true);
                // Get the top stack frame
                var frame = st.GetFrame(0);
                // Get the line number from the stack frame
                var line = frame?.GetFileLineNumber();
                return  StatusCode(500, ex.Message);
            }
        }
    }
}