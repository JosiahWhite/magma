/*
 * Licensed to the OpenAirInterface (OAI) Software Alliance under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The OpenAirInterface Software Alliance licenses this file to You under
 * the terms found in the LICENSE file in the root of this source tree.
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *-------------------------------------------------------------------------------
 * For more information about the OpenAirInterface (OAI) Software Alliance:
 *      contact@openairinterface.org
 */

#include <stdint.h>

#include "lte/gateway/c/core/common/common_defs.h"
#include "lte/gateway/c/core/oai/common/TLVDecoder.h"
#include "lte/gateway/c/core/oai/common/TLVEncoder.h"
#include "lte/gateway/c/core/oai/tasks/nas/emm/msg/AttachReject.h"

int decode_attach_reject(attach_reject_msg* attach_reject, uint8_t* buffer,
                         uint32_t len) {
  uint32_t decoded = 0;
  int decoded_result = 0;

  // Check if we got a NULL pointer and if buffer length is >= minimum length
  // expected for the message.
  CHECK_PDU_POINTER_AND_LENGTH_DECODER(buffer, ATTACH_REJECT_MINIMUM_LENGTH,
                                       len);

  /*
   * Decoding mandatory fields
   */
  if ((decoded_result = decode_emm_cause(&attach_reject->emmcause, 0,
                                         buffer + decoded, len - decoded)) < 0)
    return decoded_result;
  else
    decoded += decoded_result;

  /*
   * Decoding optional fields
   */
  while (len > decoded) {
    uint8_t ieiDecoded = *(buffer + decoded);

    /*
     * Type | value iei are below 0x80 so just return the first 4 bits
     */
    if (ieiDecoded >= 0x80) ieiDecoded = ieiDecoded & 0xf0;

    switch (ieiDecoded) {
      case ATTACH_REJECT_ESM_MESSAGE_CONTAINER_IEI:
        if ((decoded_result = decode_esm_message_container(
                 &attach_reject->esmmessagecontainer,
                 ATTACH_REJECT_ESM_MESSAGE_CONTAINER_IEI, buffer + decoded,
                 len - decoded)) <= 0)
          return decoded_result;

        decoded += decoded_result;
        /*
         * Set corresponding mask to 1 in presencemask
         */
        attach_reject->presencemask |=
            ATTACH_REJECT_ESM_MESSAGE_CONTAINER_PRESENT;
        break;

      default:
        errorCodeDecoder = TLV_UNEXPECTED_IEI;
        return TLV_UNEXPECTED_IEI;
    }
  }

  return decoded;
}

int encode_attach_reject(attach_reject_msg* attach_reject, uint8_t* buffer,
                         uint32_t len) {
  int encoded = 0;
  int encode_result = 0;

  /*
   * Checking IEI and pointer
   */
  CHECK_PDU_POINTER_AND_LENGTH_ENCODER(buffer, ATTACH_REJECT_MINIMUM_LENGTH,
                                       len);

  if ((encode_result = encode_emm_cause(&attach_reject->emmcause, 0,
                                        buffer + encoded, len - encoded)) <
      0)  // Return in case of error
    return encode_result;
  else
    encoded += encode_result;

  if ((attach_reject->presencemask &
       ATTACH_REJECT_ESM_MESSAGE_CONTAINER_PRESENT) ==
      ATTACH_REJECT_ESM_MESSAGE_CONTAINER_PRESENT) {
    if ((encode_result = encode_esm_message_container(
             attach_reject->esmmessagecontainer,
             ATTACH_REJECT_ESM_MESSAGE_CONTAINER_IEI, buffer + encoded,
             len - encoded)) < 0)
      // Return in case of error
      return encode_result;
    else
      encoded += encode_result;
  }

  return encoded;
}
