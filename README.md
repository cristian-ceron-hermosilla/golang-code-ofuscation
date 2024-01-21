# golang-code-ofuscation
Ofuscación de código fuente Golang para ambiente Linux, Unix y Mac.

Próximamente, agregaré las correciones del build para Windows. Sin embargo, si estás construyendo algún producto o solución de software que pretendas distribuir y comercializar, te sugiero instalar una maquina Linux en VirtualBox e investigar como aumentar al máximo la seguridad de la VM como ambiente de desarrollo (firewall, servicios y usuarios innecesarios, encriptar el disco, etc).  

> **Atención:** Antes de compilar y ejecutar los build, por favor revisa y trata de leer los programas y los batch.
(Si no te sientes comodo o seguro, no los utilices)

## source-ofus
Esta carpeta contiene una utilidad que permite "ofuscar" el código fuente Golang.

Este proceso de ofuscación básicamente consiste en modificar los nombres de variables y funciones en el código fuente, con el fin de aumentar la dificultad de lectura y análisis de tu codigo, en caso de caer en manos "No autorizadas".

## go-source-dic
Esta carpeta contiene el código de un programa auxiliar, cuyo objetivo es agregar al diccionario "reservedWords.txt" los comandos, funciones y/o métodos de las diversas librerías utilitarias del lenguaje para que sean IGNORADAS ya que, esto generará errores de compilaciòn.

## Como trabaja el programa "go-source-ofus"
Una vez compilado con "_build.sh" (linux/unix/mac) o "_build.bat" (windows), el programa "go-source-ofus" modifica TODOS los fuentes .go que se encuenten en tu carpeta de trabajo u "origen" en una carpeta de compilación o "destino".

     go-source-ofus ./proyecto-de-trajajo <./proyecto-ofuscado>

#### Ejemplo de Código Fuente Normal 
```go
var names = make(map[string]string)
var reservedWords []string

func main() {
	if len(os.Args) != 3 {
		println("Usage: go run main.go <source_folder> <destination_folder>")
		return
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Error al obtener el directorio de inicio del usuario: %v\n", err)
		return
	}

	sourceFolder := os.Args[1]
	destinationFolder := os.Args[2]

	err = filepath.Walk(sourceFolder, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(sourceFolder, path)
		if err != nil {
			return err
		}

		destPath := filepath.Join(destinationFolder, relPath)

		if info.IsDir() {
			return os.MkdirAll(destPath, info.Mode())
		}
```
#### Luego de la Ofuscación
```go
var nUlUx_wo_03h_anaJ8ys4_soObC_vgO8_k__ncnHT34TNeFlahgZboFGi1M = make(map[string]string)
var B_Z__5Q_A6mH9C_706__k_zE1__P_uXr_Zd__oL__n__n_dA_M_4p_Grrwv []string

func main() {
	if len(os.Args) != 3 {
		println("Usage: go run main.go <source_folder> <destination_folder>")
		return
	}

	PKwvqFMuaOUhI_uA____8d0_tmfX_cCDb_c__ZhG_H_m_q_VIUVKnRjUESr, m__BSs6k8evp8i_R71_laaa9Bm6c5_u_3J44X0F_I_IYt9DlbPmmPG_ltf_ := os.UserHomeDir()
	if m__BSs6k8evp8i_R71_laaa9Bm6c5_u_3J44X0F_I_IYt9DlbPmmPG_ltf_ != nil {
		fmt.Printf("Error al obtener el directorio de inicio del usuario: %v\n", m__BSs6k8evp8i_R71_laaa9Bm6c5_u_3J44X0F_I_IYt9DlbPmmPG_ltf_)
		return
	}

	_c_SRJVwitFTNYB6KfCe___t_0vMW7__y_mt9l____MZVYjnU__rTz1pmZ9 := os.Args[1]
	N6OlGpssT_M0YB6bPQ9Ke_jHF_W7ZysV6_XUcHzOkSRb2Hse7_UE_rmHtkJ := os.Args[2]

	m__BSs6k8evp8i_R71_laaa9Bm6c5_u_3J44X0F_I_IYt9DlbPmmPG_ltf_ = filepath.Walk(_c_SRJVwitFTNYB6KfCe___t_0vMW7__y_mt9l____MZVYjnU__rTz1pmZ9, func(path string, info fs.FileInfo, m__BSs6k8evp8i_R71_laaa9Bm6c5_u_3J44X0F_I_IYt9DlbPmmPG_ltf_ error) error {
		if m__BSs6k8evp8i_R71_laaa9Bm6c5_u_3J44X0F_I_IYt9DlbPmmPG_ltf_ != nil {
			return m__BSs6k8evp8i_R71_laaa9Bm6c5_u_3J44X0F_I_IYt9DlbPmmPG_ltf_
		}

		relPath, m__BSs6k8evp8i_R71_laaa9Bm6c5_u_3J44X0F_I_IYt9DlbPmmPG_ltf_ := filepath.Rel(_c_SRJVwitFTNYB6KfCe___t_0vMW7__y_mt9l____MZVYjnU__rTz1pmZ9, path)
		if m__BSs6k8evp8i_R71_laaa9Bm6c5_u_3J44X0F_I_IYt9DlbPmmPG_ltf_ != nil {
			return m__BSs6k8evp8i_R71_laaa9Bm6c5_u_3J44X0F_I_IYt9DlbPmmPG_ltf_
		}

		qP2_uSy___PCfzcJepVzA6M__UlkqsksP9m__ukFz_jpn__HRe0Goe2JX_R := filepath.Join(N6OlGpssT_M0YB6bPQ9Ke_jHF_W7ZysV6_XUcHzOkSRb2Hse7_UE_rmHtkJ, relPath)

		if info.IsDir() {
			return os.MkdirAll(qP2_uSy___PCfzcJepVzA6M__UlkqsksP9m__ukFz_jpn__HRe0Goe2JX_R, info.Mode())
		}

		if strings.HasSuffix(path, ".go") {
			uul_hNl__oC_HsVywzarpxk___6VjUqHOxY_kSqJNnkvoSd6Zb___CI_Zoj(PKwvqFMuaOUhI_uA____8d0_tmfX_cCDb_c__ZhG_H_m_q_VIUVKnRjUESr, path, qP2_uSy___PCfzcJepVzA6M__UlkqsksP9m__ukFz_jpn__HRe0Goe2JX_R)
			return nil
		}

		return copyFile(path, qP2_uSy___PCfzcJepVzA6M__UlkqsksP9m__ukFz_jpn__HRe0Goe2JX_R)
	})

	if m__BSs6k8evp8i_R71_laaa9Bm6c5_u_3J44X0F_I_IYt9DlbPmmPG_ltf_ != nil {
		println("Error:", m__BSs6k8evp8i_R71_laaa9Bm6c5_u_3J44X0F_I_IYt9DlbPmmPG_ltf_.Error())
	}
```
 ¿Agregadable no?!!!  :-) 
 
Si no colocas el 2do. parametro (una carpeta destino) o, por error, colocas la misma carpeta en ambos parametros se generará una carpeta "proyecto-de-trajajo-ofus".

Como asumimos que, si tu proyecto es hiper-ultra secreto, como todo proyecto que pretende comercializarse; Imaginamos que no pretenderas subirlo a Github en un repo privado ni mucho menos publico (yo tampoco lo haría), "go-source-ofus" antes de hacer lo que debe hacer, realizará un respaldo de tus fuentes en la carpeta ./backup/<fecha>/proyecto-de-trabajo/<timestamp>.

Te sugiero mantener SIEMPRE tu VM desconectada de la red e internet, salvo en los casos en que requieras actualizar o agregar dependencias a tu proyecto, tomando todos los resguardos necesarios con tu codigo y te sugiero utilizar fuentes ofuscados para agregar dependencias.

## ¿Al agregar tantas letras y simbolos, el binario será mas grande?
No. El proceso de compilaciòn traduce los identificadores a direcciones de memoria.

## Mantención del diccionario
Al compilarse los binarios, idealmente con el batch "_build" que corresponda, el diccionario "reservedWords.txt" se moverá a la carpeta del usuario $HOME (linux/unix/mac) o %USERPROFILE% (windows).

Es posible que el programa "go-source-dic" o "go-source-dic.exe" segun corresponda, agregue al diccionario identificadores (variables o funciones) que no deban omitirse en el proceso de ofuscación; Puedes quitar del diccionario los identificadores que no desees omitir.

## ¿Porque DEBERÍA ofuscar mi código en go, sí al compilarlo se genera un binario en código solo legible por la máquina?

DEBERÍAS OFUSCARLO sí estas construyendo o distrubuyendo una solución que pretendes comercializar ya que (uno) Es TU PROPIEDAD INTELECTUAL y debes EXTREMAR las medidas de protección y (dos), es totalmente factible incorporar en el binario tu codigo fuente al momento de la compilaciòn. De hecho, este es el comando:

          go build -a -buildmode=pie

## Si, esta bien. Mi código fuente podría incorporarse en el binario pero, solo si yo se lo indico "explícitamente".

¿Estas 200% seguro?.



