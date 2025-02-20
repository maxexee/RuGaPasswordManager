using System;
using System.Text;
using System.Security.Claims;
using Microsoft.IdentityModel.Tokens;
using System.IdentityModel.Tokens.Jwt;

using RuGa.Data.Identities;
using RuGaPasswordManager.Interfaces.Account_Interfaces;

namespace RuGaPasswordManager.Services
{
    public class TokenService : ITokenService
    {
        private readonly    IConfiguration  _config;
        private readonly    SymmetricSecurityKey    _key;

        public TokenService(IConfiguration  config)
        {
            _config =   config;
            _key    =   new SymmetricSecurityKey(Encoding.UTF8.GetBytes(_config["JWT:SignInKey"]));
        }
        public string CreateToken(AppUser   user)
        {
            // CREACION DEL *ClaimPrincipal".
            var claims  =   new List<Claim>{
                new Claim(JwtRegisteredClaimNames.Email,    user.Email),
                new Claim(JwtRegisteredClaimNames.GivenName,    user.UserName)
            };

            // DEFINICION DEL ALGORITMO DE ENCRIPTACION A USAR.
            var encryptionType  =   new SigningCredentials(_key,    SecurityAlgorithms.HmacSha512Signature);

            // 
            var tokenDescriptor =   new SecurityTokenDescriptor {
                Subject =   new ClaimsIdentity(claims),
                Expires =   DateTime.Now.AddDays(1),
                SigningCredentials  =   encryptionType,
                Issuer  =   _config["JWT:Issuer"],
                Audience=   _config["JWT:Audience"]
            };

            // CREACION DEL TOKEN.
            var tokenHandler    =   new JwtSecurityTokenHandler();
            var token   =   tokenHandler.CreateToken(tokenDescriptor);
            return  tokenHandler.WriteToken(token);// RETORNA EL TOKEN COMO STRING.
        }
    }
}